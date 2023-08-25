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
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	storefrontUtil "github.com/MyriadFlow/storefront-gateway/util/pkg/storefront"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	}
}

func NewStorefront(c *gin.Context) {
	var StorefrontRequest StorefrontRequest
	walletAddress := c.GetString("walletAddress")
	if err := c.BindJSON(&StorefrontRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := storefrontUtil.CreateStorefront(StorefrontRequest.Name, StorefrontRequest.Owner, walletAddress, StorefrontRequest.Plan, StorefrontRequest.Cost, StorefrontRequest.Currency, StorefrontRequest.CreatedBy, StorefrontRequest.UpdatedBy, StorefrontRequest.Image, StorefrontRequest.Headline, StorefrontRequest.Description, StorefrontRequest.Blockchain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Storefront created successfully"})
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
	date, err := time.Parse("2006-01-02", updateRequest.Validity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	storefront.Status = updateRequest.Status
	storefront.Validity = date
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
			Name:    contract.ContractName + strconv.Itoa(i),
			Address: contract.ContractAddress,
		})
	}

	graphReqBody := GraphRequest{
		Name:      req.Name,
		Folder:    req.StorefrontId,
		NodeURL:   req.NodeUrl + ":8020",
		IpfsURL:   req.NodeUrl + ":5001",
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
		StorefrontName: req.Name,
	}

	nodectlReqBytes, err := json.Marshal(nodectlReqBody)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	nodectlReq, err := http.NewRequest(http.MethodPost, req.NodectlUrl+"/marketplace", bytes.NewReader(nodectlReqBytes))
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
	defer nodectlResp.Body.Close()

	nodectlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return
	}
	fmt.Println(string(nodectlBody))
	// var nodectlRespBody NodectlResponse
	// if err := json.Unmarshal(nodectlBody, &nodectlRespBody); err != nil {
	// 	logrus.Error(err)
	// 	httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
	// 	return
	// }
	fmt.Printf("client: response body: %s\n", nodectlBody)
	httphelper.SuccessResponse(c, "succesfully deployed storefront", nil)
}
