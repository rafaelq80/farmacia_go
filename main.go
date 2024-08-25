package main

import (
	_ "github.com/rafaelq80/farmacia_go/docs"
	"github.com/rafaelq80/farmacia_go/server"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title						E-commerce - Farmácia
// @version						1.0
// @description					Projeto E-commerce - Farmácia
// @contact.name				Rafael Queiróz
// @contact.email				rafaelproinfo@gmail.com
// @contact.url					https://github.com/rafaelq80
// @license.name				Apache 2.0
// @license.url					https://www.apache.org/licenses/LICENSE-2.0.html
// @schemes						https
// @host						farmacia-go.onrender.com
// @BasePath					/
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {

	// Cria o Servidor
	app := server.SetupServer("remoto", false)

	//Rota Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Inicializa o Servidor
	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}

}
