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

// Servidor Local
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
		data.ConnectDB(config.AppConfig.ConnectionString, profile, false)
	case "teste":
		data.ConnectDB(config.AppConfig.ConnectionString, "teste", drop)
	case "remoto":
		data.ConnectDB(config.AppConfig.ConnectionString, profile, false)
	default:
		log.Println("Perfil inválido!")
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

// Servidor Local
// func SetupServer() *fiber.App{

// 	// Carregar as Congigurações do Ambiente
// 	config.LoadAppConfig("config")

// 	// Conectar com o banco de dados
// 	data.ConnectDB(config.AppConfig.ConnectionString, "local", false)

// 	app := fiber.New()

// 	// Inicializar o Log
// 	app.Use(logger.New())

// 	// Configurar o CORS
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "*",
// 		AllowHeaders: "*",
// 		AllowMethods: "*",
// 	}))

// 	// Definir as Rotas (Endpoints)
// 	routes.SetRotas(app)

// 	return app
// }

// // Servidor Remoto
// func SetupRenderServer() *fiber.App{

// 	// Carregar as Congigurações do Ambiente
// 	config.LoadAppConfig("/etc/secrets/secrets")

// 	// Conectar com o banco de dados
// 	data.ConnectDB(config.AppConfig.ConnectionString, "remoto", false)

// 	app := fiber.New()

// 	// Inicializar o Log
// 	app.Use(logger.New())

// 	// Configurar o CORS
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "*",
// 		AllowHeaders: "*",
// 		AllowMethods: "*",
// 	}))

// 	// Definir as Rotas (Endpoints)
// 	routes.SetRotas(app)

// 	return app
// }

// //Servidor de Testes
// func SetupTestServer(drop bool) *fiber.App {

// 	app := fiber.New()

// 	// Carregar as Congigurações do Ambiente
// 	config.LoadAppConfig("config")

// 	// Conectar com o banco de dados
// 	data.ConnectDB(config.AppConfig.ConnectionString, "teste", drop)

// 	// Definir as Rotas (Endpoints)
// 	routes.SetRotas(app)

// 	return app
// }
