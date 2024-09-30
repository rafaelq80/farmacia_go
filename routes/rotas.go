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
	roleService := service.NewRoleService()
	middleware := middleware.NewAuthMiddleware(usuarioService)

	// Injeção de Dependências - Classes Controladoras
	produtoController := controller.NewProdutoController(produtoService, categoriaService)
	categoriaController := controller.NewCategoriaController(categoriaService)
	usuarioController := controller.NewUsuarioController(usuarioService)
	roleController := controller.NewRoleController(roleService)

	// Rotas do Recurso Produto

	app.Route("/produtos", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), produtoController.FindAll)
		router.Post("", middleware.AuthMiddleware(RoleAdmin), produtoController.Create)
		router.Put("", middleware.AuthMiddleware(RoleAdmin), produtoController.Update)
	})

	app.Route("/produtos/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), produtoController.FindById)
		router.Delete("", middleware.AuthMiddleware(RoleAdmin), produtoController.Delete)
	})

	app.Route("/produtos/nome/:nome", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), produtoController.FindByNome)
	})

	// Rotas do Recurso Categoria

	app.Route("/categorias", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), categoriaController.FindAll)
		router.Post("", middleware.AuthMiddleware(RoleAdmin), categoriaController.Create)
		router.Put("", middleware.AuthMiddleware(RoleAdmin), categoriaController.Update)
	})

	app.Route("/categorias/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), categoriaController.FindById)
		router.Delete("", middleware.AuthMiddleware(RoleAdmin), categoriaController.Delete)
	})

	app.Route("/categorias/grupo/:grupo", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), categoriaController.FindByGrupo)
	})

	// Rotas do Recurso Usuário

	// Cadastrar Usuário (Não Protegida)
	app.Route("/usuarios/cadastrar", func(router fiber.Router) {
		router.Post("", usuarioController.Create)
	})

	// Autenticar Usuário - Login (Não Protegida)
	app.Route("/usuarios/logar", func(router fiber.Router) {
		router.Post("", usuarioController.Autenticar)
	})

	app.Route("/usuarios/all", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin), usuarioController.FindAll)
	})

	app.Route("/usuarios/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin, RoleUser), usuarioController.FindById)
	})

	app.Route("/usuarios/atualizar", func(router fiber.Router) {
		router.Put("", middleware.AuthMiddleware(RoleAdmin, RoleUser), usuarioController.Update)
	})

	// Rotas do Recurso Role

	app.Route("/roles", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin), roleController.FindAll)
		router.Post("", middleware.AuthMiddleware(RoleAdmin), roleController.Create)
		router.Put("", middleware.AuthMiddleware(RoleAdmin), roleController.Update)
	})

	app.Route("/roles/:id", func(router fiber.Router) {
		router.Get("", middleware.AuthMiddleware(RoleAdmin), roleController.FindById)
	})
}
