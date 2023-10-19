package storefront

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/constants/blockchains"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/storefront")
	{
		g.Use(paseto.PASETO)
		g.POST("", NewStorefront)
		g.PUT("", UpdateStorefront)
		g.GET("", GetStorefronts)
		g.GET("/testnet", GetTestnetStorefronts)
		g.GET("/mainnet", GetMainnetStorefronts)
		g.GET("/myStorefronts", GetStorefrontsByAddress)
		g.POST("/deploy", DeployStorefront)
		g.GET("/get_storefront_by_id", GetStorefrontById)
		g.GET("/deployment", GetDeployment)
		g.POST("/deployment/update/:id", UpdateStorefrontWebapp)
		g.POST("/deployment/stop/:id", StopStorefrontWebapp)
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
	var chain blockchains.Blockchain

	if StorefrontRequest.Network == "testnet" {
		chain = blockchains.Testnets[StorefrontRequest.Blockchain]
	} else if StorefrontRequest.Network == "mainnet" {
		chain = blockchains.Mainnets[StorefrontRequest.Blockchain]
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "network invalid"})
		return
	}

	db := dbconfig.GetDb()
	result := db.Model(&models.Storefront{}).Where("name = ?", StorefrontRequest.Name).Find(&models.Storefront{})
	if result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Storefront with the name already exists"})
		return
	}

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
	storefront.Deployed = false
	result = db.Create(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	contractReqBody := `
	{
		"data": {
			"contractName" : "AccessMaster",
    		"constructorParams":{
				"param1" : "` + walletAddress + `"
    		}
		},
		"network" : "` + chain.DeploymentName + `"
	}
	`

	contractReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.SMARTCONTRACT_API_URL+"/api/Contract", strings.NewReader(contractReqBody))
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
	storefront.StorefrontHeadline = updateRequest.Headline
	storefront.StorefrontDescription = updateRequest.Description
	storefront.ProfileImage = updateRequest.ProfileImage
	storefront.StorefrontImage = updateRequest.StorefrontImage
	storefront.PersonalTagline = updateRequest.PersonalTagline
	storefront.PersonalDescription = updateRequest.PersonalDescription
	storefront.RelevantImage = updateRequest.RelevantImage
	storefront.MailId = updateRequest.MailId
	storefront.Twitter = updateRequest.Twitter
	storefront.Discord = updateRequest.Discord
	storefront.Instagram = updateRequest.Instagram
	storefront.UpdatedAt = time.Now()
	storefront.Region = updateRequest.Region
	storefront.Type = updateRequest.Type
	storefront.Category = updateRequest.Category
	storefront.Tags = updateRequest.Tags
	result = db.Save(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error, "message": "error in updating storefront in database"})
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
	err := db.Model(&models.Storefront{}).Where("wallet_address = ?", walletAddress).Find(&storefronts).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Storefronts fetched successfully", storefronts)
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
	result := db.Model(&models.Storefront{}).Where("id = ?", req.Id).First(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	var chain blockchains.Blockchain

	if storefront.Network == "testnet" {
		chain = blockchains.Testnets[storefront.Blockchain]
	} else if storefront.Network == "mainnet" {
		chain = blockchains.Mainnets[storefront.Blockchain]
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "network invalid"})
		return
	}
	storefront.Name = req.Name
	storefront.StorefrontHeadline = req.Headline
	storefront.StorefrontDescription = req.Description
	storefront.ProfileImage = req.ProfileImage
	storefront.StorefrontImage = req.StorefrontImage
	storefront.PersonalTagline = req.PersonalTagline
	storefront.PersonalDescription = req.PersonalDescription
	storefront.RelevantImage = req.RelevantImage
	storefront.MailId = req.MailId
	storefront.Twitter = req.Twitter
	storefront.Discord = req.Discord
	storefront.Instagram = req.Instagram
	storefront.UpdatedAt = time.Now()
	storefront.Region = req.Region
	storefront.Type = req.Type
	storefront.Category = req.Category
	storefront.Tags = req.Tags

	var contracts []models.Contract
	err := db.Model(&models.Contract{}).Where("storefront_id = ?", req.Id).Find(&contracts).Error
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
	subgraphNameReq := strings.ReplaceAll(req.Name, " ", "")
	graphName := req.Tag + "/" + subgraphNameReq
	graphReqBody := GraphRequest{
		Name:      graphName,
		Folder:    req.Id.String(),
		NodeURL:   envconfig.EnvVars.SUBGRAPH_SERVER_URL + ":" + chain.GraphPort,
		IpfsURL:   envconfig.EnvVars.SUBGRAPH_SERVER_URL + ":" + chain.IpfsPort,
		Contracts: reqContracts,
		Network:   chain.SubgraphNetworkName,
		Protocol:  "ethereum",
	}

	graphReqBytes, err := json.Marshal(graphReqBody)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	fmt.Println(string(graphReqBytes))
	graphReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.SMARTCONTRACT_API_URL+"/api/Subgraph", bytes.NewReader(graphReqBytes))
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
	subgraphUrl := chain.GraphHttpsUrl + "/subgraphs/name/" + graphName + "/graphql"
	subgraphId := subgraphIdArr[2]
	subgraph = models.Subgraph{
		SubgraphId:    subgraphId,
		Name:          graphName,
		Network:       chain.SubgraphNetworkName,
		Protocol:      "ethereum",
		SubgraphUrl:   subgraphUrl,
		WalletAddress: walletAddress,
		StorefrontId:  req.Id.String(),
	}

	db.Create(&subgraph)
	domain := strings.ReplaceAll(req.Name, " ", "")
	nodectlReqBody := NodectlRequest{
		StorefrontName: domain,
		StorefrontId:   req.Id.String(),
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
	storefront.Deployed = true

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

func GetDeployment(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var res []GetDeploymentPayload
	err := db.Model(&models.Storefront{}).Where("wallet_address = ?", walletAddress).Find(&res).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "fetched data successfully", res)
}

func GetTestnetStorefronts(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var storefronts []models.Storefront
	result := db.Model(&models.Storefront{}).Where("wallet_address = ? AND network = ?", walletAddress, "testnet").Find(&storefronts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, storefronts)
}

func GetMainnetStorefronts(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var storefronts []models.Storefront
	result := db.Model(&models.Storefront{}).Where("wallet_address = ? AND network = ?", walletAddress, "mainnet").Find(&storefronts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, storefronts)
}

func UpdateStorefrontWebapp(c *gin.Context) {
	storefrontId := c.Param("id")
	resp, err := http.Get(envconfig.EnvVars.NODECTL_SERVER_URL + ":" + envconfig.EnvVars.NODECTL_SERVER_PORT + "/marketplace/update/" + storefrontId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated marketplace"})
}

func StopStorefrontWebapp(c *gin.Context) {
	storefrontId := c.Param("id")
	resp, err := http.Get(envconfig.EnvVars.NODECTL_SERVER_URL + ":" + envconfig.EnvVars.NODECTL_SERVER_PORT + "/marketplace/stop/" + storefrontId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()
	c.JSON(http.StatusOK, gin.H{"message": "Successfully stopped marketplace"})
}
