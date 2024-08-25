package main

import (
	"github.com/rafaelq80/farmacia_go/server"
)

//	@title						E-commerce - Farmácia
//	@version					1.0
//	@description				Projeto E-commerce - Farmácia
//	@contact.name				Rafael Queiróz
//	@contact.email				rafaelproinfo@gmail.com
//	@contact.url				https://github.com/rafaelq80
//	@license.name				Apache 2.0
//	@license.url				https://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:8000
//	@BasePath					/
//	@schemes					http
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func main() {

	// Cria o Servidor Local
	//app := server.SetupServer()
	
	// Cria o Servidor Remoto
	app := server.SetupRenderServer()
	
	// Inicializa o Servidor
	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}

}
