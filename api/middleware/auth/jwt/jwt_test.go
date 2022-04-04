package jwt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TheLazarusNetwork/marketplace-engine/config"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/models"
	"github.com/TheLazarusNetwork/marketplace-engine/models/claims"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/auth"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"
	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_JWT(t *testing.T) {
	config.Init("../../../../.env")
	logwrapper.Init("../../../../logs")
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
		token, err := auth.GenerateToken(newClaims, envutil.MustGetEnv("JWT_PRIVATE_KEY"))
		if err != nil {
			t.Fatal(err)
		}
		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusOK, statusCode)
	})

	t.Run("Should return 403 with incorret JWT", func(t *testing.T) {
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateToken(newClaims, "this private key is valid key")
		if err != nil {
			t.Fatal(err)
		}
		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusForbidden, statusCode)
	})

	t.Run("Should return 403 with expired JWT", func(t *testing.T) {
		expiration := time.Now().Add(time.Second * 2)
		signedBy := envutil.MustGetEnv("SIGNED_BY")
		newClaims := claims.CustomClaims{
			WalletAddress: testWalletAddress,
			SignedBy:      signedBy,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiration),
			},
		}
		time.Sleep(time.Second * 2)
		token, err := auth.GenerateToken(newClaims, envutil.MustGetEnv("JWT_PRIVATE_KEY"))
		if err != nil {
			t.Fatal(err)
		}

		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusForbidden, statusCode)
	})

}

func callApi(t *testing.T, token string) int {
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
	return rr.Result().StatusCode
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
