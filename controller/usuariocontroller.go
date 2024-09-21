package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
	security "github.com/rafaelq80/farmacia_go/security/bcrypt"
	"github.com/rafaelq80/farmacia_go/service"
	"github.com/rafaelq80/farmacia_go/validator"
)

// Injeção de Dependências - UsuarioService
type UsuarioController struct {
	service *service.UsuarioService
}

// Método Construtor
func NewUsuarioController(service *service.UsuarioService) *UsuarioController {
	return &UsuarioController{service: service}
}

//	@Summary		Listar Usuarios
//	@Description	Lista todos os Usuarios
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{array}		model.Usuario
//	@Failure		401				{object}	config.HTTPError
//	@Router			/usuarios/all 	[get]
//	@Security		Bearer
func (usuarioController *UsuarioController) FindAllUsuario(context *fiber.Ctx) error {

	usuarios := usuarioController.service.FindAll()

	return context.Status(fiber.StatusOK).JSON(usuarios)

}

//	@Summary		Listar Usuario por id
//	@Description	Lista um Usuario por id
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Id do Usuario"
//	@Success		200				{array}		model.Usuario
//	@Failure		400				{object}	config.HTTPError
//	@Failure		401				{object}	config.HTTPError
//	@Failure		404				{object}	config.HTTPError
//	@Router			/usuarios/{id} 	[get]
//	@Security		Bearer
func (usuarioController *UsuarioController) FindByIdUsuario(context *fiber.Ctx) error {

	id := context.Params("id")

	usuario, found := usuarioController.service.FindById(id)

	if !found {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
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
//	@Failure		400			{object}	config.HTTPError
//	@Failure		404			{object}	config.HTTPError
//	@Router			/usuarios 	[post]
func (usuarioController *UsuarioController) CreateUsuario(context *fiber.Ctx) error {

	var usuario model.Usuario

	if err := context.BodyParser(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	// Verifica se o usuário existe (evitar duplicidade)
	if usuarioController.service.EmailExists(usuario.Usuario) {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Usuário já cadastrado!"})
	}

	if err := usuarioController.service.Create(&usuario); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error creating usuario"})
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
//	@Failure		400						{object}	config.HTTPError
//	@Failure		401						{object}	config.HTTPError
//	@Failure		404						{object}	config.HTTPError
//	@Router			/usuarios/atualizar 	[put]
//	@Security		Bearer
func (usuarioController *UsuarioController) UpdateUsuario(context *fiber.Ctx) error {

	var usuario model.Usuario

	if err := context.BodyParser(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&usuario); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	var id = strconv.FormatUint(uint64(usuario.ID), 10)

	// Verifica se o Usuário existe
	if !usuarioController.service.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Usuário não encontrado!"})
	}

	// Localiza os dados do Usuário
	buscarUsuario := usuarioController.service.FindByUsuario(usuario.Usuario)

	// Verifica se o e-mail já pertence a outro usuário
	if usuarioController.service.EmailExists(usuario.Usuario) && usuario.ID != buscarUsuario.ID {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Usuário já cadastrado!"})
	}

	// Criptografa a senha
	senhaCriptografada, _ := security.HashPassword(usuario.Senha)
	usuario.Senha = senhaCriptografada

	data.DB.Save(&usuario)

	return context.Status(fiber.StatusOK).JSON(&usuario)
}

//	@Summary		Autenticar Usuario
//	@Description	Autentica um Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			usuario				body		model.UsuarioLogin	true	"Autenticar Usuario"
//	@Success		200					{object}	model.UsuarioLogin
//	@Failure		400					{object}	config.HTTPError
//	@Failure		401					{object}	config.HTTPError
//	@Failure		404					{object}	config.HTTPError
//	@Router			/usuarios/logar 						[post]
func (usuarioController *UsuarioController) AutenticarUsuario(c *fiber.Ctx) error {

	var usuarioLogin model.UsuarioLogin

	if err := c.BodyParser(&usuarioLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	usuarioAutenticado, err := usuarioController.service.AutenticarUsuario(&usuarioLogin)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(usuarioAutenticado)

}
