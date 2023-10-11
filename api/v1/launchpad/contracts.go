package launchpad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MyriadFlow/storefront-gateway/config/constants/blockchains"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/gin-gonic/gin"
)

func Deploy(c *gin.Context, link string) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var req reqBody
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//get blockcahin from storefront
	var storefront models.Storefront
	err = db.Model(&models.Storefront{}).Where("storefront_id = ?", req.StorefrontId).First(&storefront).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//check for invalid actions
	if req.ContractName == "TradeHub" {
		result := db.Model(&models.Contract{}).Where("contract_name = ? AND storefront_id = ?", req.ContractName, req.StorefrontId).Find(&models.Contract{})
		if result.RowsAffected != 0 {
			//tradehub exists
			c.JSON(http.StatusBadRequest, gin.H{"message": "Tradehub already deployed"})
			return
		}
	}

	if req.ContractName != "TradeHub" && req.ContractName != "AccessMaster" {
		result := db.Model(&models.Contract{}).Where("contract_name = ? AND storefront_id = ?", "TradeHub", req.StorefrontId).Find(&models.Contract{})
		if result.RowsAffected == 0 {
			//tradehub not deployed
			c.JSON(http.StatusBadRequest, gin.H{"message": "Tradehub not Deployed"})
			return
		}
	}

	contractReqBody := contractReqBody{
		Data: data{
			ContractName:      req.ContractName,
			ConstructorParams: req.ConstructorParams,
		},
		Network: blockchains.Testnets[storefront.Blockchain].DeploymentName,
	}
	contractReqBodyBytes, err := json.Marshal(contractReqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	contractReq, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(contractReqBodyBytes))
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
	strn = string(body)[1 : len(strn)-1]
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
		ContractName:    req.ContractName,
		ContractAddress: response.ContractAddress,
		WalletAddress:   walletAddress,
		ChainId:         response.ChainId,
		Verified:        response.Verified,
		StorefrontId:    req.StorefrontId,
		BlockNumber:     response.BlockNumber,
		CollectionName:  req.CollectionName,
		Thumbnail:       req.Thumbnail,
		CoverImage:      req.CoverImage,
	}
	result := db.Create(&contract)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeployContract(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/api/Contract", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}
