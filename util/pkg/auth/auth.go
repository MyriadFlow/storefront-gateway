package auth

import (
	"github.com/vk-rv/pvx"
)

func GenerateTokenPaseto(claims pvx.Claims, privateKey string) (string, error) {

	ask := pvx.NewAsymmetricSecretKey([]byte(privateKey), pvx.Version4)

	pv4 := pvx.NewPV4Public()

	token, err := pv4.Sign(ask, claims)

	if err != nil {
		return "", err
	}
	return token, nil
}
