package delegateartifactcreation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MyriadFlow/storefront-gateway/global"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"math/big"
)

func TestDelegateAssetCreation(t *testing.T) {
	var sampleRoyalty *big.Int
	time.Sleep(4 * time.Second)
	testingcommon.InitializeEnvVars()
	logwrapper.Init()
	global.InitGlobal()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWallet := testingcommon.GenerateWallet()
	createrWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	url := "/api/v1.0/delegateAssetCreation"
	sampleRoyalty=big.NewInt(500)
	t.Run("Should fail if wallet address is not hexadecimal", func(t *testing.T) {
		rr := httptest.NewRecorder()
		reqBody := DelegateAssetCreationRequest{
			CreatorAddress: "invalidwalletaddr",
			MetaDataHash:   "ipfs://QmTiQKxZoVMvDahqVUzvkJhAjF9C1MzytpDEocxUT3oBde",
			RoyaltyPercentBasisPoint : sampleRoyalty,
		}
		jsonBytes, _ := json.Marshal(reqBody)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		deletegateAssetCreation(c)
		ok := assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode, rr.Body.String())
		if !ok {
			t.FailNow()
		}
	})
	t.Run("Should be able to delegate asset", func(t *testing.T) {
		rr := httptest.NewRecorder()
		reqBody := DelegateAssetCreationRequest{
			CreatorAddress: createrWallet.WalletAddress,
			MetaDataHash:   "ipfs://QmTiQKxZoVMvDahqVUzvkJhAjF9C1MzytpDEocxUT3oBde",
			RoyaltyPercentBasisPoint : sampleRoyalty,
		}
		jsonBytes, _ := json.Marshal(reqBody)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		deletegateAssetCreation(c)
		ok := assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
		if !ok {
			t.FailNow()
		}
	})

}
