package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/swaggo/files" // swagger embed files
)

// HTTPError representa um erro HTTP genérico
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Error message"`
}

// RegisterSwaggerErrors registra as rotas do Swagger na aplicação
func RegisterSwaggerErrors(app *fiber.App) {
	// Registrar as rotas Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // A rota que serve a documentação Swagger
}

// getErrorMessage retorna a mensagem de erro apropriada para o status HTTP
func getErrorMessage(status int) string {
	switch status {
	case fiber.StatusBadRequest:
		return "Bad Request"
	case fiber.StatusUnauthorized:
		return "Unauthorized"
	case fiber.StatusForbidden:
		return "Forbidden"
	case fiber.StatusNotFound:
		return "Not Found"
	default:
		return "Internal Server Error"
	}
}

// SwaggerMiddleware é um middleware para lidar com erros e registrar documentação automaticamente
func SwaggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Executa o próximo handler
		err := c.Next()

		// Se houver um erro, adiciona a documentação do Swagger
		if err != nil {
			status := c.Response().StatusCode()
			message := getErrorMessage(status)
			addSwaggerErrorDoc(c.Path(), c.Method(), status, message)
		}

		return err
	}
}

// addSwaggerErrorDoc adiciona a documentação de erro ao Swagger
func addSwaggerErrorDoc(path, method string, status int, message string) {
	// Adiciona a documentação do erro no console ou outra lógica de log/documentação
	fmt.Printf("Erro documentado: Path: %s, Method: %s, Status: %d, Message: %s\n", path, method, status, message)
}
