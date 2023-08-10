package subgraph

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subgraph")
	{
		g.POST("", DeploySubgraph)
	}
}

func DeploySubgraph(c *gin.Context) {
	var req SubgraphPayload
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

	body, _ := io.ReadAll(resp.Body)
	strn := string(body)
	strn = string(body)[1 : len(strn)-1]
	data, err := base64.StdEncoding.DecodeString(strn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	arr := strings.Split(string(data), "\n")
	subgraphId := strings.Split(arr[3], " ")[2]
	subgraphUrl := strings.Split(arr[5], " ")[2]
	res := SubgraphResponse{
		SubgraphUrl: subgraphUrl,
		SubgraphId:  subgraphId,
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}
