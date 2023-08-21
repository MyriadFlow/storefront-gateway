package subgraph

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subgraph")
	{
		g.Use(paseto.PASETO)
		g.POST("", DeploySubgraph)
		g.GET("", GetSubgraphsByAddress)
	}
}

func DeploySubgraph(c *gin.Context) {
	var req SubgraphPayload
	walletAddress := c.GetString("walletAddress")
	if err := c.BindJSON(&req); err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	deployReq, err := http.NewRequest(http.MethodPost, envconfig.EnvVars.SMARTCONTRACT_API_URL+"/Subgraph", bytes.NewReader(reqBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(deployReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	db := dbconfig.GetDb()
	var subgraph models.Subgraph
	body, _ := io.ReadAll(resp.Body)
	strn := string(body)
	strn = string(body)[1 : len(strn)-1]
	data, err := base64.StdEncoding.DecodeString(strn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	arr := strings.Split(string(data), "\n")
	subgraphIdArr := strings.Split(arr[3], " ")
	subgraphUrlArr := strings.Split(arr[5], " ")
	subgraphUrl := subgraphUrlArr[2]
	subgraphId := subgraphIdArr[2]
	subgraph = models.Subgraph{
		SubgraphId:      subgraphId,
		Name:            req.Name,
		Folder:          req.Folder,
		NodeUrl:         req.NodeUrl,
		IpfsUrl:         req.IpfsUrl,
		ContractAddress: req.ContractAddress,
		Network:         req.Network,
		Protocol:        req.Protocol,
		Tag:             req.Tag,
		SubgraphUrl:     subgraphUrl,
		WalletAddress:   walletAddress,
	}

	db.Create(&subgraph)

	res := SubgraphResponse{
		SubgraphUrl: subgraph.SubgraphUrl,
		SubgraphId:  subgraph.SubgraphId,
	}
	c.JSON(http.StatusOK, res)
}

func GetSubgraphsByAddress(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var subgraphs []models.Subgraph
	err := db.Model(&models.Subgraph{}).Where("wallet_address = ?", walletAddress).Find(&subgraphs)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", subgraphs)
}
