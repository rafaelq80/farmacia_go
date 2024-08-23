package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rafaelq80/farmacia_go/config"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/routes"

)

// Servidor Principal
func SetupServer() *fiber.App{

	// Carregar as Congigurações do Ambiente
	config.LoadAppConfig()

	// Conectar com o banco de dados
	data.ConnectDB(config.AppConfig.ConnectionString)

	app := fiber.New()

	// Inicializar o Log
	app.Use(logger.New())

	// Configurar o CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	// Definir as Rotas (Endpoints)
	routes.SetRotas(app)

	return app
}

//Servidor de Testes
func SetupTestServer(dropTables bool) *fiber.App {
	
	app := fiber.New()

	// Carregar as Congigurações do Ambiente
	config.LoadAppConfig()

	// Conectar com o banco de dados
	data.ConnectTestDB(config.AppConfig.TestConnectionString, dropTables)

	// Definir as Rotas (Endpoints)
	routes.SetRotas(app)

	return app
}
