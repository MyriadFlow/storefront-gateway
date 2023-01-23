package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	gopaseto "aidanwoods.dev/go-paseto"
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/MyriadFlow/storefront_gateway/models/claims"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"
)

var PublicKey gopaseto.V4AsymmetricPublicKey
var secretKey gopaseto.V4AsymmetricSecretKey

func Init() {
	secretKey = gopaseto.NewV4AsymmetricSecretKey()
	PublicKey = secretKey.Public()
}
func GenerateTokenPaseto(claim claims.CustomClaims, privateKey string) (string, error) {
	footer := envconfig.EnvVars.FOOTER
	claimbyte, _ := json.Marshal(claim)
	fmt.Println("claim value", claimbyte)
	token, err := gopaseto.NewTokenFromClaimsJSON(claimbyte, []byte(footer))
	if err != nil {
		return "", err
	}
	pasetoExpirationInHours, ok := os.LookupEnv("PASETO_EXPIRATION_IN_HOURS")
	pasetoExpirationInHoursInt := time.Duration(24)
	if ok {
		res, err := strconv.Atoi(pasetoExpirationInHours)
		if err != nil {
			logwrapper.Log.Warnf("Failed to parse PASETO_EXPIRATION_IN_HOURS as int : %v", err.Error())
		} else {
			pasetoExpirationInHoursInt = time.Duration(res)
		}
	}
	pasetoExpirationHours := pasetoExpirationInHoursInt * time.Hour
	expiration := time.Now().Add(pasetoExpirationHours)
	token.SetExpiration(expiration)
	signed := token.V4Sign(secretKey, nil)
	return signed, nil
}

func Getpublickey() gopaseto.V4AsymmetricPublicKey {
	publickey := PublicKey
	return publickey
}
