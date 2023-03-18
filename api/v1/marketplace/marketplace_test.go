package marketplace

import (
	"net/http"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	// "github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	// "github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	// "github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_PatchMarketplaceById(t *testing.T){
	testingcommon.InitializeEnvVars()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	//url := "/api/v1.0/marketplace"

	t.Run("Should be able to update token id", func(t *testing.T) {

		rr := httptest.NewRecorder()
		itemId:=testingcommon.CreateTestMarketplace(t)

		url := "/api/v1.0/marketplace/"+string(itemId)

		fmt.Println("url hit : ",url)
		requestBody := models.Marketplace{
			TokenId: "111111",
		}
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)

		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)

		patchMarketplaceById(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}