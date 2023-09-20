package storefront

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/storefront")
	{
		g.Use(paseto.PASETO)
		g.POST("", NewStorefront)
		g.PUT("", UpdateStorefront)
		g.GET("", GetStorefronts)
		g.GET("/myStorefronts", GetStorefrontsByAddress)
		g.POST("/deploy", DeployStorefront)
		g.GET("/get_storefront_by_id", GetStorefrontById)
	}
}

type resBody struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	Verified        bool   `json:"verified"`
	BlockNumber     int    `json:"blockNumber"`
}

func NewStorefront(c *gin.Context) {
	var StorefrontRequest StorefrontRequest
	walletAddress := c.GetString("walletAddress")
	if err := c.BindJSON(&StorefrontRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	db := dbconfig.GetDb()
	err := db.Model(&models.Storefront{}).Where("name = ?", StorefrontRequest.Name).Error
	if !(errors.Is(err, gorm.ErrRecordNotFound)) {
		//storefront name exists
		c.JSON(http.StatusBadRequest, gin.H{"message": "Storefront with the name already exists"})
		return
	} else {
		//storefront name doesnt exist
		id := uuid.New()
		storefront := models.Storefront{
			Id:            id,
			Name:          StorefrontRequest.Name,
			WalletAddress: walletAddress,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			Image:         StorefrontRequest.Image,
			Headline:      StorefrontRequest.Headline,
			Description:   StorefrontRequest.Description,
			Blockchain:    StorefrontRequest.Blockchain,
			Network:       StorefrontRequest.Network,
		}
		result := db.Create(&storefront)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}

		contractReqBody := `
	{
		"data": {
			"contractName" : "AccessMaster",
    		"constructorParams":{
    		}
		},
		"network" : "maticmum"
	}
	`

		contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.SMARTCONTRACT_API_URL+"/Contract", strings.NewReader(contractReqBody))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		client := &http.Client{}
		resp, err := client.Do(contractReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		strn := string(body)
		strn = string(body)[1 : len(
			strn)-1]
		data, err := base64.StdEncoding.DecodeString(strn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		arr := strings.Split(string(data), "\n")
		response := new(resBody)

		if err := json.Unmarshal([]byte(arr[len(arr)-3]), response); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		contract := models.Contract{
			ContractName:    "AccessMaster",
			ContractAddress: response.ContractAddress,
			WalletAddress:   walletAddress,
			ChainId:         response.ChainId,
			Verified:        response.Verified,
			StorefrontId:    id.String(),
			BlockNumber:     response.BlockNumber,
		}
		result = db.Create(&contract)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Storefront created successfully", "storefrontId": id, "accessMasterAddress": response.ContractAddress})
}

func UpdateStorefront(c *gin.Context) {
	db := dbconfig.GetDb()
	var updateRequest UpdateStorefrontRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var storefront models.Storefront
	result := db.Where("id = ?", updateRequest.Id).First(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	storefront.Name = updateRequest.Name
	storefront.Headline = updateRequest.Headline
	storefront.Description = updateRequest.Description
	storefront.Image = updateRequest.Image
	storefront.ProfileImage = updateRequest.ProfileImage
	storefront.CoverImage = updateRequest.CoverImage
	storefront.AssetName = updateRequest.AssetName
	storefront.AssetDescription = updateRequest.AssetDescription
	storefront.PersonalInformation = updateRequest.PersonalInformation
	storefront.PersonalDescription = updateRequest.PersonalDescription
	storefront.RelevantImage = updateRequest.RelevantImage
	storefront.MailId = updateRequest.MailId
	storefront.Twitter = updateRequest.Twitter
	storefront.Discord = updateRequest.Discord
	storefront.Instagram = updateRequest.Instagram
	storefront.UpdatedBy = updateRequest.UpdatedBy
	storefront.UpdatedAt = time.Now()
	result = db.Save(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Storefront updated successfully"})
}

func GetStorefronts(c *gin.Context) {
	db := dbconfig.GetDb()
	var storefronts []models.Storefront
	result := db.Find(&storefronts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, storefronts)
}

func GetStorefrontsByAddress(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var storefronts []models.Storefront
	err := db.Model(&models.Storefront{}).Where("wallet_address = ?", walletAddress).Find(&storefronts)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", storefronts)
}

func DeployStorefront(c *gin.Context) {
	var req DeployStorefrontRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")

	var storefront models.Storefront
	result := db.Where("id = ?", req.StorefrontId).First(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	storefront.Name = req.StorefrontName
	storefront.Headline = req.Headline
	storefront.Description = req.Description
	storefront.ProfileImage = req.ProfileImage
	storefront.CoverImage = req.CoverImage
	storefront.AssetName = req.AssetName
	storefront.AssetDescription = req.AssetDescription
	storefront.PersonalInformation = req.PersonalInformation
	storefront.PersonalDescription = req.PersonalDescription
	storefront.RelevantImage = req.RelevantImage
	storefront.MailId = req.MailId
	storefront.Twitter = req.Twitter
	storefront.Discord = req.Discord
	storefront.Instagram = req.Instagram

	var contracts []models.Contract
	err := db.Model(&models.Contract{}).Where("storefront_id = ?", req.StorefrontId).Find(&contracts).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	var reqContracts []Contract
	for i, contract := range contracts {
		reqContracts = append(reqContracts, Contract{
			Name:        contract.ContractName + strconv.Itoa(i),
			Address:     contract.ContractAddress,
			BlockNumber: contract.BlockNumber,
		})
	}

	graphReqBody := GraphRequest{
		Name:      req.Name,
		Folder:    req.StorefrontId,
		NodeURL:   envconfig.EnvVars.SUBGRAPH_SERVER_URL + ":8020",
		IpfsURL:   envconfig.EnvVars.SUBGRAPH_SERVER_URL + ":5001",
		Contracts: reqContracts,
		Network:   req.Network,
		Protocol:  req.Protocol,
		Tag:       req.Tag,
	}

	graphReqBytes, err := json.Marshal(graphReqBody)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	fmt.Println(string(graphReqBytes))
	graphReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.SMARTCONTRACT_API_URL+"/Subgraph", bytes.NewReader(graphReqBytes))
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(graphReq)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Status})
		return
	}
	defer resp.Body.Close()

	var subgraph models.Subgraph
	body, _ := io.ReadAll(resp.Body)
	strn := string(body)
	strn = string(body)[1 : len(strn)-1]
	data, err := base64.StdEncoding.DecodeString(strn)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	arr := strings.Split(string(data), "\n")
	subgraphIdArr := strings.Split(arr[3], " ")
	subgraphUrlArr := strings.Split(arr[5], " ")
	subgraphUrl := subgraphUrlArr[2]
	subgraphId := subgraphIdArr[2]
	subgraph = models.Subgraph{
		SubgraphId:    subgraphId,
		Name:          req.Name,
		Network:       req.Network,
		Protocol:      req.Protocol,
		Tag:           req.Tag,
		SubgraphUrl:   subgraphUrl,
		WalletAddress: walletAddress,
		StorefrontId:  req.StorefrontId,
	}

	db.Create(&subgraph)

	nodectlReqBody := NodectlRequest{
		StorefrontName: req.StorefrontName,
		StorefrontId:   req.StorefrontId,
	}

	nodectlReqBytes, err := json.Marshal(nodectlReqBody)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	nodectlReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.NODECTL_SERVER_URL+":"+envconfig.EnvVars.NODECTL_SERVER_PORT+"/marketplace", bytes.NewReader(nodectlReqBytes))
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	nodectlResp, err := http.DefaultClient.Do(nodectlReq)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if nodectlResp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Status})
		return
	}
	defer nodectlResp.Body.Close()

	nodectlBody, err := io.ReadAll(nodectlResp.Body)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	var nodectlRespBody NodectlResponse
	if err := json.Unmarshal(nodectlBody, &nodectlRespBody); err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	storefront.WebappUrl = nodectlRespBody.StorefrontUrl
	storefront.SubgraphUrl = subgraphUrl

	result = db.Save(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	httphelper.SuccessResponse(c, "succesfully deployed storefront", gin.H{
		"graphUrl":      subgraph.SubgraphUrl,
		"storefrontUrl": storefront.WebappUrl,
	})
}

func GetStorefrontById(c *gin.Context) {
	db := dbconfig.GetDb()
	id := c.Query("id")
	fmt.Println(id)
	if id == "" {
		logrus.Error(fmt.Errorf("%s : Failed to get the id ", "GetStorefrontsById"))
	}
	var storefront models.Storefront
	err := db.Model(&models.Storefront{}).Where("id = ?", id).Find(&storefront).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "GetStorefrontsById : Failed to get storefront")
		return
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", storefront)
}
