package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rafaelq80/farmacia_go/config"
)

func CreateToken(usuario string) (string, error) {

	config.LoadAppConfig("config")

	claims := jwt.MapClaims{}
	claims["sub"] = usuario
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.Secret))

}
