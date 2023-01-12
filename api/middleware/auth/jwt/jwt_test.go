package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MyriadFlow/storefront_gateway/api/types"

	customstatuscodes "github.com/MyriadFlow/storefront_gateway/config/constants/http/custom_status_codes"
	"github.com/MyriadFlow/storefront_gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront_gateway/config/dbconfig/dbinit"
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/MyriadFlow/storefront_gateway/models"
	"github.com/MyriadFlow/storefront_gateway/models/claims"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/auth"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"
	"github.com/MyriadFlow/storefront_gateway/util/testingcommon"
	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_JWT(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
	dbinit.Init()
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
	t.Run("Should return 200 with correct JWT", func(t *testing.T) {
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateToken(newClaims, envconfig.EnvVars.JWT_PRIVATE_KEY)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

	t.Run("Should return 401 with incorret JWT", func(t *testing.T) {
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateToken(newClaims, "this private key is valid key")
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
	})

	t.Run("Should return 401 and 4011 with expired JWT", func(t *testing.T) {
		expiration := time.Now().Add(time.Second * 2)
		signedBy := envconfig.EnvVars.SIGNED_BY
		newClaims := claims.CustomClaims{
			WalletAddress: testWalletAddress,
			SignedBy:      signedBy,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiration),
			},
		}
		time.Sleep(time.Second * 2)
		token, err := auth.GenerateToken(newClaims, envconfig.EnvVars.JWT_PRIVATE_KEY)
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

	header := fmt.Sprintf("Bearer %v", token)
	rq, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rq.Header.Add("Authorization", header)
	ginTestApp.Use(JWT)
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
