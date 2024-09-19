package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/controller"
	"github.com/rafaelq80/farmacia_go/security/middleware"
	"github.com/rafaelq80/farmacia_go/service"
)

func SetRotas(app *fiber.App) {

	// Rota de checagem do status do servidor
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "Successo!",
			"message": "Bem vindo ao Golang!",
		})
	})

	// Injeção de Dependências - Produto
	produtoService := service.NewProdutoService()
	produtoController := controller.NewProdutoController(produtoService)

	// Injeção de Dependências - Categoria
	categoriaService := service.NewCategoriaService()
	categoriaController := controller.NewCategoriaController(categoriaService)

	// Injeção de Dependências - Categoria
	usuarioService := service.NewUsuarioService()
	usuarioController := controller.NewUsuarioController(usuarioService)

	// Rotas do Recurso Produto

	app.Route("/produtos", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, produtoController.FindAllProduto)
		router.Post("", middleware.AuthMiddleware, produtoController.CreateProduto)
		router.Put("", middleware.AuthMiddleware, produtoController.UpdateProduto)
	})

	app.Route("/produtos/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, produtoController.FindByIdProduto)
		router.Delete("", middleware.AuthMiddleware, produtoController.DeleteProduto)
	})

	app.Route("/produtos/nome/:nome", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, produtoController.FindByNomeProduto)
	})

	// Rotas do Recurso Categoria

	app.Route("/categorias", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, categoriaController.FindAllCategoria)
		router.Post("", middleware.AuthMiddleware, categoriaController.CreateCategoria)
		router.Put("", middleware.AuthMiddleware, categoriaController.UpdateCategoria)
	})

	app.Route("/categorias/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, categoriaController.FindByIdCategoria)
		router.Delete("", middleware.AuthMiddleware, categoriaController.DeleteCategoria)
	})

	app.Route("/categorias/grupo/:grupo", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, categoriaController.FindByGrupoCategoria)
	})

	// Rotas do Recurso Usuário

	// Cadastrar Usuário (Não Protegida)
	app.Route("/usuarios/cadastrar", func(router fiber.Router) {
		router.Post("", usuarioController.CreateUsuario)
	})

	// Autenticar Usuário - Login (Não Protegida)
	app.Route("/usuarios/logar", func(router fiber.Router) {
		router.Post("", usuarioController.AutenticarUsuario)
	})

	app.Route("/usuarios/all", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, usuarioController.FindAllUsuario)
	})

	app.Route("/usuarios/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware, usuarioController.FindByIdUsuario)
	})

	app.Route("/usuarios/atualizar", func(router fiber.Router) {
		router.Put("", middleware.AuthMiddleware, usuarioController.UpdateUsuario)
	})

}
