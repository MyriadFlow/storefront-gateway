package profile

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MyriadFlow/storefront-gateway/api/types"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig/dbinit"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_PatchProfile(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
	dbinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	url := "/api/v1.0/profile"

	// TODO: Write more tests
	t.Run("Should be able to update name", func(t *testing.T) {

		rr := httptest.NewRecorder()

		requestBody := GetProfilePayload{
			Name: "Yash",
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

		patchProfile(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})
}

func Test_GetProfile(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	t.Run("Should be able to get profile details", func(t *testing.T) {

		testWallet := testingcommon.GenerateWallet()
		header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
		url := "/api/v1.0/profile"
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		getProfile(c)
		var response types.ApiResponse
		body := rr.Body
		json.NewDecoder(body).Decode(&response)
		var user models.User
		testingcommon.ExtractPayload(&response, &user)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		assert.Equal(t, "Jack", user.Name)
		assert.Equal(t, "https://revoticengineering.com/", user.ProfilePicture)
		assert.Equal(t, "India", user.Location)
		logrus.Debug(user)
	})

}
