package highlights

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetHighlightsItemIds(t *testing.T) {
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	t.Run("Should be able get itemIds ", func(t *testing.T) {

		rr := httptest.NewRecorder()
		url := "/api/v1.0/highlights"
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)

		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		getHighlightsItemIds(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}

func Test_DeleteHighlightsItemId(t *testing.T) {
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	t.Run("Should be able delete itemId ", func(t *testing.T) {

		rr := httptest.NewRecorder()
		itemId := testingcommon.CreateTestHighlights(t)
		url := "/api/v1.0/highlights"
		req, err := http.NewRequest("DELETE", url, nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)

		c.Request = req
		c.Params = []gin.Param{
			{
				Key:   "itemId",
				Value: itemId,
			},
		}
		c.Set("walletAddress", testWallet.WalletAddress)
		deleteHighlightsItemId(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}

func Test_PostHighlightsItemId(t *testing.T) {
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	t.Run("Should be able post item details ", func(t *testing.T) {
		rr := httptest.NewRecorder()
		testTokenId := "11111"
		url := "/api/v1.0/highlights"
		requestBody := models.Highlights{
			NFT_Contract_Address: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			TokenId:              testTokenId,
			MetaDataURI:          "hjhjgh",
		}
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)

		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		postHighlightsItemId(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		testingcommon.DeleteTestHighlights(t, testTokenId)
	})

}
