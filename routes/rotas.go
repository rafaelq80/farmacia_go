package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/controller"
	"github.com/rafaelq80/farmacia_go/security/middleware"
	
)

func SetRotas(app *fiber.App) {

	// Rota de checagem do status do servidor
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "Successo!",
			"message": "Bem vindo ao Golang!",
		})
	})

	// app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Rotas do Recurso Produto

	app.Route("/produtos", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindAllProduto)
		router.Post("", middleware.AuthMiddleware, controller.CreateProduto)
		router.Put("", middleware.AuthMiddleware, controller.UpdateProduto)
	})

	app.Route("/produtos/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindByIdProduto)
		router.Delete("", middleware.AuthMiddleware, controller.DeleteProduto)
	})

	app.Route("/produtos/nome/:nome", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindByNomeProduto)
	})

	// Rotas do Recurso Categoria

	app.Route("/categorias", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindAllCategoria)
		router.Post("", middleware.AuthMiddleware, controller.CreateCategoria)
		router.Put("", middleware.AuthMiddleware, controller.UpdateCategoria)
	})

	app.Route("/categorias/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindByIdCategoria)
		router.Delete("", middleware.AuthMiddleware, controller.DeleteCategoria)
	})

	app.Route("/categorias/grupo/:grupo", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindByGrupoCategoria)
	})

	// Rotas do Recurso Usuário

	// Cadastrar Usuário (Não Protegida)
	app.Route("/usuarios/cadastrar", func(router fiber.Router) {
		router.Post("", controller.CreateUsuario)
	})

	// Rota de Login (Não Protegida)
	app.Route("/usuarios/logar", func(router fiber.Router) {
		router.Post("", controller.AutenticarUsuario)
	})

	app.Route("/usuarios/all", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindAllUsuario)
	})

	app.Route("/usuarios/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, controller.FindByIdUsuario)
	})

	app.Route("/usuarios/atualizar", func(router fiber.Router) {
		router.Put("", middleware.AuthMiddleware, controller.UpdateUsuario)
	})

}
