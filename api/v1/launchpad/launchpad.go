package launchpad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/launchpad")
	{
		g.POST("/AccessMaster", DeployAccessMaster)
		g.POST("/TradeHub", DeployTradeHub)
		g.POST("/FusionSeries", DeployFusionSeries)
		g.POST("/SignatureSeries", DeploySignatureSeries)
		g.POST("/InstaGen", DeployInstaGen)
		g.POST("/EternumPass", DeployEternumPass)
	}
}

type res struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	Verified        bool   `json:"verified"`
}

func Deploy(c *gin.Context, link string) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	strn := string(body)
	strn = string(body)[1 : len(strn)-1]
	data, err := base64.StdEncoding.DecodeString(strn)
	if err != nil {
		log.Fatal("error:", err)
	}
	arr := strings.Split(string(data), "\n")
	//fmt.Println(arr[len(arr)-3])
	response := new(res)
	if err := json.Unmarshal([]byte(arr[len(arr)-3]), response); err != nil {
		log.Fatal("error unmarshaling")
	}
	c.JSON(http.StatusOK, response)
}

func DeployAccessMaster(c *gin.Context) {
	Deploy(c, "http://localhost:8080/AccessMaster")
}

func DeployTradeHub(c *gin.Context) {
	Deploy(c, "http://localhost:8080/TradeHub")
}

func DeployFusionSeries(c *gin.Context) {
	Deploy(c, "http://localhost:8080/FusionSeries")
}
func DeploySignatureSeries(c *gin.Context) {
	Deploy(c, "http://localhost:8080/SignatureSeries")
}

func DeployInstaGen(c *gin.Context) {
	Deploy(c, "http://localhost:8080/InstaGen")
}

func DeployEternumPass(c *gin.Context) {
	Deploy(c, "http://localhost:8080/InstaGen")
}
