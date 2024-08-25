package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/rafaelq80/farmacia_go/docs"
	"github.com/rafaelq80/farmacia_go/server"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title						E-commerce - Farmácia
// @version					1.0
// @description				Projeto E-commerce - Farmácia
// @contact.name				Rafael Queiróz
// @contact.email				rafaelproinfo@gmail.com
// @contact.url				https://github.com/rafaelq80
// @license.name				Apache 2.0
// @license.url				https://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8000
// @BasePath					/
// @schemes					http
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {

	// Cria o Servidor Local
	//app := server.SetupServer()

	// Cria o Servidor Remoto
	app := server.SetupRenderServer()

	//Rota Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/swagger/*", func(c *fiber.Ctx) error {
		host := c.Hostname() // Captura o hostname dinâmico
		swaggerURL := "/swagger/index.html?url=" + host + "/swagger/doc.json"
		return c.Redirect(swaggerURL)
	})

	// Inicializa o Servidor
	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}

}
