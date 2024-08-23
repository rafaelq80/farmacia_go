package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {

	// Extrai o Token do Header
	tokenString := c.Get("Authorization")

	// Verifica se o Token está vazio
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "401", 
			"message": "Acesso Negado!",
		})
	}

	// Valida o Token JWT (observe que retira a palavra Bearer)
	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret")), nil
	})

	// Se o token for nulo ou inválido
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "401", 
			"message": "Acesso Negado!",
		})
	}

	// Extrai as claims do token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "500", 
			"message": "Token Inválido!",
		})
	} 

	if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "401", 
			"message": "O Token Expirou!",
		})
	}

	return c.Next()

}
