package launchpad

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/launchpad")
	{
		g.POST("/FlowAccessControl", DeployFlowAccessControl)
		g.POST("/Tradehub", DeployTradeHub)
		g.POST("/FusionSeries", DeployFusionSeries)
		g.POST("/SignatureSeries", DeploySignatureSeries)
		g.POST("/InstaGen", DeployInstaGen)
	}
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

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	c.JSON(http.StatusOK, string(body))
}

func DeployFlowAccessControl(c *gin.Context) {
	Deploy(c, "http://localhost:8080/FlowAccessControl")
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
