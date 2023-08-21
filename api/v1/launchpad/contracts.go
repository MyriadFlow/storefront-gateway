package launchpad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/gin-gonic/gin"
)

type reqBody struct {
	ContractName      string         `json:"contractName"`
	ConstructorParams map[string]any `json:"constructorParams"`
	WalletAdress      string         `json:"walletAdress"`
	Network           string         `json:"network"`
	StorefrontId      string         `json:"storefrontId"`
}
type resBody struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	Verified        bool   `json:"verified"`
}
type data struct {
	ContractName      string         `json:"contractName"`
	ConstructorParams map[string]any `json:"constructorParams"`
}
type contractReqBody struct {
	Data    data   `json:"data"`
	Network string `json:"network"`
}

func Deploy(c *gin.Context, link string) {
	db := dbconfig.GetDb()
	var req reqBody
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	contractReqBody := contractReqBody{
		Data: data{
			ContractName:      req.ContractName,
			ConstructorParams: req.ConstructorParams,
		},
		Network: req.Network,
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
		WalletAdress:    req.WalletAdress,
		ChainId:         response.ChainId,
		Verified:        response.Verified,
		StorefrontId:    req.StorefrontId,
	}
	result := db.Create(&contract)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeployContract(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/Contract", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}
