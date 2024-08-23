package auth

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(usuario string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["usuario"] = usuario
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("secret")))

}
