package delegateartifactcreation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig/dbinit"
	"github.com/TheLazarusNetwork/marketplace-engine/global"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDelegateArtifactCreation(t *testing.T) {
	time.Sleep(4 * time.Second)
	logwrapper.Init()
	dbinit.Init()
	global.InitGlobal()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWallet := testingcommon.GenerateWallet()
	createrWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	url := "/api/v1.0/delegateArtifactCreation"
	t.Run("Should fail if wallet address is not hexadecimal", func(t *testing.T) {
		rr := httptest.NewRecorder()
		reqBody := DelegateArtifactCreationRequest{
			CreatorAddress: "invalidwalletaddr",
			MetaDataHash:   "ipfs://QmTiQKxZoVMvDahqVUzvkJhAjF9C1MzytpDEocxUT3oBde",
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
		deletegateArtifactCreation(c)
		ok := assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode, rr.Body.String())
		if !ok {
			t.FailNow()
		}
	})
	t.Run("Should be able to delegate artifact", func(t *testing.T) {
		rr := httptest.NewRecorder()
		reqBody := DelegateArtifactCreationRequest{
			CreatorAddress: createrWallet.WalletAddress,
			MetaDataHash:   "ipfs://QmTiQKxZoVMvDahqVUzvkJhAjF9C1MzytpDEocxUT3oBde",
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
		deletegateArtifactCreation(c)
		ok := assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
		if !ok {
			t.FailNow()
		}
	})

}
