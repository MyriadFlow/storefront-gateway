package testingcommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"os"

	"github.com/MyriadFlow/storefront-gateway/api/types"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/models/claims"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/auth"

	"crypto/ecdsa"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
)

func PrepareAndGetAuthHeader(t *testing.T, testWalletAddress string) string {
	gin.SetMode(gin.TestMode)
	CreateTestUser(t, testWalletAddress)
	customClaims := claims.New(testWalletAddress)
	token, err := auth.GenerateTokenPaseto(customClaims)
	if err != nil {
		t.Fatal(err)
	}
	header := fmt.Sprintf("Bearer %v", token)

	return header
}

func CreateTestUser(t *testing.T, walletAddress string) {
	db := dbconfig.GetDb()
	user := models.User{
		Name:              "Jack",
		ProfilePictureUrl: "https://testdomain.com/",
		WalletAddress:     walletAddress,
		Country:           "India",
	}
	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		t.Fatal(err)
	}
}

func CreateTestHighlights(t *testing.T) string {
	db := dbconfig.GetDb()
	highlights := models.Highlights{
		NFT_Contract_Address: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		TokenId:              "1234",
		MetaDataURI:          "hjhjgh",
	}
	err := db.Model(&models.Highlights{}).Create(&highlights).Error
	if err != nil {
		t.Fatal(err)
	}

	var itemId int
	err = db.Model(&models.Highlights{}).Select("item_id").Where("nft_contract_address = ? AND token_id = ? AND meta_data_uri = ? ", highlights.NFT_Contract_Address, highlights.TokenId, highlights.MetaDataURI).First(&itemId).Error
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%v", itemId)
}

func DeleteTestHighlights(t *testing.T, tokenId string) {
	db := dbconfig.GetDb()

	err := db.Where("token_id = ?", tokenId).Delete(&models.Highlights{}).Error
	if err != nil {
		t.Fatal(err)
	}

}

func GenerateWallet() *TestWallet {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	pvtkey := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	testWallet := TestWallet{
		PrivateKey:    pvtkey,
		WalletAddress: address,
	}
	return &testWallet
}

// Converts map created by json decoder to struct
// out should be pointer (&payload)
func ExtractPayload(response *types.ApiResponse, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(response.Payload)
	json.NewDecoder(buf).Decode(out)
}

func InitializeEnvVars() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	envconfig.EnvVars.APP_NAME = "storefront-gateway"
	envconfig.EnvVars.APP_PORT = 8000
	envconfig.EnvVars.APP_MODE = "debug"
	envconfig.EnvVars.APP_ALLOWED_ORIGIN = []string{"*"}

	envconfig.EnvVars.DB_HOST = hostname
	envconfig.EnvVars.DB_USERNAME = "postgres"
	envconfig.EnvVars.DB_PASSWORD = "root"
	envconfig.EnvVars.DB_NAME = "lazarus"
	envconfig.EnvVars.DB_PORT = 5432

	envconfig.EnvVars.AUTH_EULA = "I Accept the MyriadFlow Terms of Service https://myriadflow.com/terms.html for accessing the application. Challenge ID: "
	envconfig.EnvVars.MARKETPLACE_CONTRACT_ADDRESS = "0x72AFc9D60EBd2265a2420d580D2918392fae47f6"
	envconfig.EnvVars.STOREFRONT_CONTRACT_ADDRESS = "0xe0CdEbF537574BcbB362885593Ee896D58Aa88Ec"
	envconfig.EnvVars.POLYGON_RPC = "https://rpc-mumbai.maticvigil.com/v1/f336dfba703440ee198bf937d5c065b8fe04891c"
	envconfig.EnvVars.MNEMONIC = "immense area chuckle ritual voyage certain script fury oil pill month affair"
	envconfig.EnvVars.NFT_STORAGE_API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDllYzI4RDAyRkFiN0Q3MjM5NTI4RjA0QjhiMkZlMEJGYzU1QTVDNDciLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTY3NzIzMzM5Nzc3OSwibmFtZSI6ImFjY2Vzcy10b2tlbiJ9.QtV7r6-VLSDyfKJlCtnFPR9gwPVRhtQwbrIpmZNPYrM"
	envconfig.EnvVars.PASETO_SIGNED_BY = "MyriadFlow"
	envconfig.EnvVars.PASETO_FOOTER = "MyriadFlow 2023"
	envconfig.EnvVars.PASETO_EXPIRATION_IN_HOURS = "168"

	envconfig.EnvVars.ALLOWED_WALLET_ADDRESS = []string{"*"}

}
