package routes

import (
	"github.com/rafaelq80/farmacia_go/controller"
	_ "github.com/rafaelq80/farmacia_go/docs"
	"github.com/rafaelq80/farmacia_go/security/middleware"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
)

func SetRotas(app *fiber.App) {

	//Rota Swagger
	app.Route("/swagger", func(router fiber.Router) {
		router.Get("*", fiberSwagger.WrapHandler)
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "Successo!",
			"message": "Bem vindo ao Golang! Fiber Ativo!",
		})
	})

	// Cadastrar Usuário
	app.Route("/usuarios/cadastrar", func(router fiber.Router) {
		router.Post("", controller.CreateUsuario)
	})

	// Rota de Login
	app.Route("/usuarios/logar", func(router fiber.Router) {
		router.Post("", controller.AutenticarUsuario)
	})

	// Rotas Protegidas
	app.Use(middleware.AuthMiddleware)

	// Rotas do Recurso Produto

	app.Route("/produtos", func(router fiber.Router) {
		router.Get("", controller.FindAllProduto)
		router.Post("", controller.CreateProduto)
		router.Put("", controller.UpdateProduto)
	})

	app.Route("/produtos/:id", func(router fiber.Router) {
		router.Get("", controller.FindByIdProduto)
		router.Delete("", controller.DeleteProduto)
	})

	app.Route("/produtos/nome/:nome", func(router fiber.Router) {
		router.Get("", controller.FindByNomeProduto)
	})

	// Rotas do Recurso Categoria

	app.Route("/categorias", func(router fiber.Router) {
		router.Get("", controller.FindAllCategoria)
		router.Post("", controller.CreateCategoria)
		router.Put("", controller.UpdateCategoria)
	})

	app.Route("/categorias/:id", func(router fiber.Router) {
		router.Get("", controller.FindByIdCategoria)
		router.Delete("", controller.DeleteCategoria)
	})

	app.Route("/categorias/grupo/:grupo", func(router fiber.Router) {
		router.Get("", controller.FindByGrupoCategoria)
	})

	// Rotas do Recurso Usuário

	app.Route("/usuarios/all", func(router fiber.Router) {
		router.Get("", controller.FindAllUsuario)
	})

	app.Route("/usuarios/:id", func(router fiber.Router) {
		router.Get("", controller.FindByIdUsuario)
	})

	app.Route("/usuarios/atualizar", func(router fiber.Router) {
		router.Put("", controller.UpdateUsuario)
	})

}
