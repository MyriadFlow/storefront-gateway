package paseto

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MyriadFlow/storefront-gateway/api/types"
	customstatuscodes "github.com/MyriadFlow/storefront-gateway/config/constants/http/custom_status_codes"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/models/claims"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/auth"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PASETO(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
	db := dbconfig.GetDb()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWalletAddress := testingcommon.GenerateWallet().WalletAddress
	newUser := models.User{
		WalletAddress: testWalletAddress,
	}
	err := db.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		db.Delete(&newUser)
	}()
	t.Run("Should return 200 with correct PASETO", func(t *testing.T) {
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateTokenPaseto(newClaims)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

	// t.Run("Should return 401 with incorret PASETO", func(t *testing.T) {
	// 	newClaims := claims.New(testWalletAddress)
	// 	token, err := auth.GenerateTokenPaseto(newClaims, "aaaabbaa")
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	rr := callApi(t, token)
	// 	assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
	// })

	t.Run("Should return 401 and 4011 with expired PASETO", func(t *testing.T) {
		expiration := time.Now().Add(time.Second * 2)
		signedBy := envconfig.EnvVars.PASETO_SIGNED_BY
		newClaims := claims.CustomClaims{
			WalletAddress: testWalletAddress,
			SignedBy:      signedBy,
			Expiration:    expiration,
		}
		time.Sleep(time.Second * 2)
		token, err := auth.GenerateTokenPaseto(newClaims)
		if err != nil {
			t.Fatal(err)
		}

		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
		var response types.ApiResponse
		body := rr.Body
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, customstatuscodes.TokenExpired, response.StatusCode)
	})

}

func callApi(t *testing.T, token string) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ginTestApp := gin.New()

	rq, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rq.Header.Add("Authorization", token)
	ginTestApp.Use(PASETO)
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
