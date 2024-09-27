package config

import (
	"github.com/gofiber/fiber/v2"
)

// Exemplo de Erro 
func NewError(c *fiber.Ctx, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(er)
}

// Exemplo de Erro HTTP
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}