package paseto

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/MyriadFlow/storefront_gateway/models/claims"
	"gorm.io/gorm"

	"github.com/vk-rv/pvx"

	customstatuscodes "github.com/MyriadFlow/storefront_gateway/config/constants/http/custom_status_codes"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func PASETO(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {
		err = fmt.Errorf("failed to bind header, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(headers.Authorization, ErrAuthHeaderMissing)
		httphelper.ErrResponse(c, http.StatusBadRequest, ErrAuthHeaderMissing.Error())
		c.Abort()
		return
	}
	pasetoToken := headers.Authorization
	fmt.Println("header token", pasetoToken)
	pv4 := pvx.NewPV4Public()
	k := envconfig.EnvVars.PASETO_PUBLIC_KEY
	symK := pvx.NewAsymmetricPublicKey([]byte(k), pvx.Version4)
	var cc claims.CustomClaims
	err = pv4.
		Verify(pasetoToken, symK).
		ScanClaims(&cc)

	if err != nil {
		var validationErr *pvx.ValidationError
		if errors.As(err, &validationErr) {
			if validationErr.HasExpiredErr() {
				err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
				logValidationFailed(headers.Authorization, err)
				httphelper.CErrResponse(c, http.StatusUnauthorized, customstatuscodes.TokenExpired, "token expired")
				c.Abort()
				return
			}

		}
		err = fmt.Errorf("failed to scan claims for paseto token, %s", err)
		logValidationFailed(headers.Authorization, err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		if err := cc.Valid(); err != nil {
			logValidationFailed(headers.Authorization, err)
			if err.Error() == gorm.ErrRecordNotFound.Error() {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				err = fmt.Errorf("failed to validate claim, %s", err)
				logwrapper.Log.Error(err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		} else {
			c.Set("walletAddress", cc.WalletAddress)
			c.Next()
		}
	}
}

func logValidationFailed(token string, err error) {
	logwrapper.Warnf("validation failed with token %v and error: %v", token, err)
}
