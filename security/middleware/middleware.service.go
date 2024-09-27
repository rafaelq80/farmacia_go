package middleware

import (
    "strings"
    "time"
    "github.com/gofiber/fiber/v2"
    jwt "github.com/golang-jwt/jwt/v5"
    "github.com/rafaelq80/farmacia_go/config"
    "github.com/rafaelq80/farmacia_go/service"
)

type AuthMiddleware struct {
    usuarioService *service.UsuarioService
}

func NewAuthMiddleware(usuarioService *service.UsuarioService) *AuthMiddleware {
    return &AuthMiddleware{
        usuarioService: usuarioService,
    }
}

func (am *AuthMiddleware) AuthMiddleware(requiredRoles ...uint) fiber.Handler {
    return func(c *fiber.Ctx) error {
        config.LoadAppConfig("config")

        token, err := am.extractAndValidateToken(c)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":  "401",
                "message": err.Error(),
            })
        }

        claims, err := am.extractAndValidateClaims(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":  "401",
                "message": err.Error(),
            })
        }

        usuario := claims["sub"].(string)
        buscaUsuario, err := am.usuarioService.FindByUsuario(usuario)
        if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "status":  "404",
                "message": "Usuário não encontrado",
            })
        }

        if err := am.checkRoles(buscaUsuario.RoleID, requiredRoles); err != nil {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "status":  "403",
                "message": err.Error(),
            })
        }

        // Set user information in context for later use if needed
        c.Locals("user", buscaUsuario)

        return c.Next()
    }
}

func (am *AuthMiddleware) extractAndValidateToken(c *fiber.Ctx) (*jwt.Token, error) {
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        return nil, fiber.NewError(fiber.StatusUnauthorized, "Token não fornecido")
    }

    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.AppConfig.Secret), nil
    })

    if err != nil || !token.Valid {
        return nil, fiber.NewError(fiber.StatusUnauthorized, "Token inválido")
    }

    return token, nil
}

func (am *AuthMiddleware) extractAndValidateClaims(token *jwt.Token) (jwt.MapClaims, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, fiber.NewError(fiber.StatusInternalServerError, "Token inválido")
    }

    if expiresAt, ok := claims["exp"]; ok {
        if int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
            return nil, fiber.NewError(fiber.StatusUnauthorized, "Token expirado")
        }
    }

    return claims, nil
}

func (am *AuthMiddleware) checkRoles(userRole uint, requiredRoles []uint) error {
    if len(requiredRoles) == 0 {
        return nil
    }

    for _, role := range requiredRoles {
        if role == userRole {
            return nil
        }
    }

    return fiber.NewError(fiber.StatusForbidden, "Você não possui o Direito de Acesso!")
}