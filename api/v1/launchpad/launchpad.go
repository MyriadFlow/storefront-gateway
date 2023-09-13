package launchpad

import (
	"context"
	"log"
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/launchpad")
	{
		g.Use(paseto.PASETO)
		g.GET("/contracts", GetContracts)
		g.GET("/contracts/:storefrontId", GetContractsById)
		g.POST("/contract", DeployContract)
		g.GET("/contracts/:storefrontId/:contractName", GetContractsByName)
		g.GET("/address", GetAddressFromHash)

	}
}
func GetContracts(c *gin.Context) {
	db := dbconfig.GetDb()
	var contracts []models.Contract
	if result := db.Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

func GetContractsById(c *gin.Context) {
	storefrontId := c.Param("storefrontId")
	walletAddress := c.GetString("walletAddress")
	db := dbconfig.GetDb()
	var contracts []models.Contract
	if result := db.Model(models.Contract{}).Where("storefront_id = ? AND wallet_address = ?", storefrontId, walletAddress).Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

func GetContractsByName(c *gin.Context) {
	contractName := c.Param("contractName")
	storefrontId := c.Param("storefrontId")
	walletAddress := c.GetString("walletAddress")
	db := dbconfig.GetDb()
	var contracts []models.Contract
	if result := db.Model(models.Contract{}).Where("contract_name = ? AND storefront_id = ? AND wallet_address = ?", contractName, storefrontId, walletAddress).Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

type AddressReq struct {
	TransactionHash string `json:"transactionHash"`
}

func GetAddressFromHash(c *gin.Context) {
	var req AddressReq

	ethereumURL := envconfig.EnvVars.POLYGON_RPC_HTTP
	tx := req.TransactionHash

	ethereumClient, err := ethclient.Dial(ethereumURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	ctx := context.Background()
	transactionHash := common.HexToHash(tx)

	transaction, _, err := ethereumClient.TransactionByHash(ctx, transactionHash)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, transaction.To())
}
