package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rafaelq80/farmacia_go/config"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/routes"
)

func SetupServer(profile string, drop bool) *fiber.App {

	// Carregar as Congigurações do Ambiente
	if profile == "remoto" {
		config.LoadAppConfig("/etc/secrets/secrets")
	} else {
		config.LoadAppConfig("config")
	}

	// Conectar com o banco de dados
	switch profile {
	case "local":
		data.ConnectDB(config.AppConfig.ConnectionString, profile, drop)
	case "remoto":
		data.ConnectDB(config.AppConfig.ConnectionString, profile, drop)
	case "teste":
		data.ConnectDB(config.AppConfig.ConnectionString, profile, drop)
	default:
		log.Println("Database Profile inválido!")
	}

	// Inicializa uma nova instância do Fiber
	app := fiber.New()

	// Inicializa o Log
	app.Use(logger.New())

	// Configura o CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	// Definir as Rotas (Endpoints)
	routes.SetRotas(app)

	return app
}
