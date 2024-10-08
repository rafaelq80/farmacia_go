package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/service"
	"github.com/rafaelq80/farmacia_go/validator"
)

// Injeção de Dependências - UsuarioService
type UsuarioController struct {
	usuarioService *service.UsuarioService
}

// Método Construtor
func NewUsuarioController(usuarioService *service.UsuarioService) *UsuarioController {
	return &UsuarioController{usuarioService: usuarioService}
}

//	@Summary		Listar Usuarios
//	@Description	Lista todos os Usuarios
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{array}	model.Usuario
//	@Router			/usuarios/all 	[get]
//	@Security		Bearer
func (usuarioController *UsuarioController) FindAll(context *fiber.Ctx) error {

	usuarios, _ := usuarioController.usuarioService.FindAll()

	return context.Status(fiber.StatusOK).JSON(usuarios)

}

//	@Summary		Listar Usuario por id
//	@Description	Lista um Usuario por id
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			id				path	string	true	"Id do Usuario"
//	@Success		200				{array}	model.Usuario
//	@Router			/usuarios/{id} 	[get]
//	@Security		Bearer
func (usuarioController *UsuarioController) FindById(context *fiber.Ctx) error {

	id := context.Params("id")

	usuario, err := usuarioController.usuarioService.FindById(id)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Usuário não encontrado!",
		})
	}

	return context.Status(fiber.StatusOK).JSON(usuario)

}

//	@Summary		Criar Usuario
//	@Description	Cria um novo Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			usuario		body		model.Usuario	true	"Criar Usuario"
//	@Success		201			{object}	model.Usuario
//	@Router			/usuarios 	[post]
func (usuarioController *UsuarioController) Create(context *fiber.Ctx) error {

	var usuario model.Usuario

	if err := context.BodyParser(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err,
		})
	}

	// Verifica se o usuário existe (evitar duplicidade)
	if usuarioController.usuarioService.EmailExists(usuario.Usuario) {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Usuário já cadastrado!",
		})
	}

	if err := usuarioController.usuarioService.Create(&usuario); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(&usuario)

}

//	@Summary		Atualizar Usuario
//	@Description	Edita os dados de um Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			Usuario					body		model.Usuario	true	"Atualizar Usuario"
//	@Success		200						{object}	model.Usuario
//	@Router			/usuarios/atualizar 	[put]
//	@Security		Bearer
func (usuarioController *UsuarioController) Update(context *fiber.Ctx) error {

	var usuario model.Usuario

	if err := context.BodyParser(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err,
		})
	}

	id := strconv.FormatUint(uint64(usuario.ID), 10)
	usuarioExist, _ := usuarioController.usuarioService.ExistsById(id)

	if !usuarioExist {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Usuário não encontrado!",
		})
	}

	// Localiza os dados do Usuário
	buscarUsuario, _ := usuarioController.usuarioService.FindByUsuario(usuario.Usuario)

	// Verifica se o e-mail já pertence a outro usuário
	if usuarioController.usuarioService.EmailExists(usuario.Usuario) && usuario.ID != buscarUsuario.ID {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "O e-mail informado já pertence a outro usuário!",
		})
	}

	// Atualiza os dados do Usuário
	if err := usuarioController.usuarioService.Update(&usuario); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&usuario)
}

//	@Summary		Autenticar Usuario
//	@Description	Autentica um Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			usuario				body		model.UsuarioLogin	true	"Autenticar Usuario"
//	@Success		200					{object}	model.UsuarioLogin
//	@Router			/usuarios/logar 																								[post]
func (usuarioController *UsuarioController) Autenticar(c *fiber.Ctx) error {

	var usuarioLogin model.UsuarioLogin

	if err := c.BodyParser(&usuarioLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err.Error(),
		})
	}

	usuarioAutenticado, err := usuarioController.usuarioService.AutenticarUsuario(&usuarioLogin)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "401",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(usuarioAutenticado)

}
