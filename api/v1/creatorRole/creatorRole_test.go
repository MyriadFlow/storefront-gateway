//go:build !race
// +build !race

package creatorrole

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	storefront "github.com/MyriadFlow/storefront-gateway/config/storefront"
	"github.com/MyriadFlow/storefront-gateway/global"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostGrantRole(t *testing.T) {
	defer time.Sleep(4 * time.Second)

	testingcommon.InitializeEnvVars()
	logwrapper.Init()
	global.InitGlobal()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)

	t.Run("Should not grant if not have operator role", func(t *testing.T) {
		testWallet := testingcommon.GenerateWallet()

		headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
		url := "/api/v1.0/creatorrole"
		rr := httptest.NewRecorder()
		//get sample test creatorRoeId
		creatorRoleId, err := storefront.GetRole(storefront.CREATOR_ROLE)
		if err != nil {
			logwrapper.Fatal(err)
		}

		reqBody := CreatorRoleRequest{
			RoleId: string(creatorRoleId[:]),
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
		postGrantRole(c)
		assert.Equal(t, http.StatusForbidden, rr.Result().StatusCode)
	})

}
