package config

import (
	"github.com/gofiber/fiber/v2"
)

func NewError(c *fiber.Ctx, status int, err error) {
	var er interface{}

	switch status {
	case 400:
		er = HTTPError400{
			Code:    status,
			Message: "Bad Request",
		}
	case 401:
		er = HTTPError401{
			Code:    status,
			Message: "Unauthorized",
		}
	case 403:
		er = HTTPError403{
			Code:    status,
			Message: "Forbidden",
		}
	case 404:
		er = HTTPError404{
			Code:    status,
			Message: "Not Found",
		}
	default:
		er = HTTPError500{
			Code:    status,
			Message: "Internal Server Error",
		}
	}

	c.Status(status).JSON(er)
}

// Exemplos de Erro HTTP para documentação Swagger

// HTTPError400 representa um erro de Bad Request
type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

// HTTPError401 representa um erro de Unauthorized
type HTTPError401 struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"Unauthorized"`
}

// HTTPError403 representa um erro de Forbidden
type HTTPError403 struct {
	Code    int    `json:"code" example:"403"`
	Message string `json:"message" example:"Forbidden"`
}

// HTTPError404 representa um erro de Not Found
type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found"`
}

// HTTPError500 representa um erro de Internal Server Error
type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}
