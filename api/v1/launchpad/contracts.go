package launchpad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/gin-gonic/gin"
)

type res struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	Verified        bool   `json:"verified"`
}

func Deploy(c *gin.Context, link string) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	req, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
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
	//fmt.Println(arr[len(arr)-3])
	response := new(res)

	if err := json.Unmarshal([]byte(arr[len(arr)-3]), response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeployAccessMaster(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/AccessMaster", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}

func DeployTradeHub(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/TradeHub", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}

func DeployFusionSeries(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/FusionSeries", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}
func DeploySignatureSeries(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/SignatureSeries", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}

func DeployInstaGen(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/InstaGen", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}

func DeployEternumPass(c *gin.Context) {
	Deploy(c, fmt.Sprintf("%s/EternumPass", envconfig.EnvVars.SMARTCONTRACT_API_URL))
}
