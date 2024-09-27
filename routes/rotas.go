package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/controller"
	"github.com/rafaelq80/farmacia_go/security/middleware"
	"github.com/rafaelq80/farmacia_go/service"
)

const (
    RoleAdmin = 1
    RoleUser  = 2
)

func SetRotas(app *fiber.App) {

	// Rota de checagem do status do servidor
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "Successo!",
			"message": "Bem vindo ao Golang!",
		})
	})

	
	// Inicialização dos serviços
	produtoService := service.NewProdutoService()
	categoriaService := service.NewCategoriaService()
	emailService := service.NewEmailService()
	usuarioService := service.NewUsuarioService(emailService)
	middleware := middleware.NewAuthMiddleware(usuarioService)

	// Injeção de Dependências - Produto
	produtoController := controller.NewProdutoController(produtoService, categoriaService)

	// Injeção de Dependências - Categoria
	categoriaController := controller.NewCategoriaController(categoriaService)

	// Injeção de Dependências - Usuário
	usuarioController := controller.NewUsuarioController(usuarioService)

	// Rotas do Recurso Produto

	app.Route("/produtos", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), produtoController.FindAllProduto)
		router.Post("", middleware.AuthMiddleware(RoleAdmin), produtoController.CreateProduto)
		router.Put("", middleware.AuthMiddleware(RoleAdmin), produtoController.UpdateProduto)
	})

	app.Route("/produtos/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), produtoController.FindByIdProduto)
		router.Delete("", middleware.AuthMiddleware(RoleAdmin), produtoController.DeleteProduto)
	})

	app.Route("/produtos/nome/:nome", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), produtoController.FindByNomeProduto)
	})

	// Rotas do Recurso Categoria

	app.Route("/categorias", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), categoriaController.FindAllCategoria)
		router.Post("", middleware.AuthMiddleware(RoleAdmin), categoriaController.CreateCategoria)
		router.Put("", middleware.AuthMiddleware(RoleAdmin), categoriaController.UpdateCategoria)
	})

	app.Route("/categorias/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), categoriaController.FindByIdCategoria)
		router.Delete("", middleware.AuthMiddleware(RoleAdmin), categoriaController.DeleteCategoria)
	})

	app.Route("/categorias/grupo/:grupo", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), categoriaController.FindByGrupoCategoria)
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
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), usuarioController.FindAllUsuario)
	})

	app.Route("/usuarios/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), usuarioController.FindByIdUsuario)
	})

	app.Route("/usuarios/atualizar", func(router fiber.Router) {
		router.Put("", middleware.AuthMiddleware(RoleAdmin), usuarioController.UpdateUsuario)
	})

}
