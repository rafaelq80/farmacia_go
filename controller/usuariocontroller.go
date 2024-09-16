package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
	security "github.com/rafaelq80/farmacia_go/security/bcrypt"
	"github.com/rafaelq80/farmacia_go/validator"
)

//	@Summary		Listar Usuarios
//	@Description	Lista todos os Usuarios
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Usuario
//	@Failure		401	{object}	config.HTTPError
//	@Router			/usuarios/all [get]
//	@Security		Bearer
func FindAllUsuario(c *fiber.Ctx) error {

	var usuarios []model.Usuario

	data.DB.Preload("Produto").Find(&usuarios)

	return c.Status(fiber.StatusOK).JSON(usuarios)
}

//	@Summary		Listar Usuario por id
//	@Description	Lista um Usuario por id
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Id do Usuario"
//	@Success		200	{array}		model.Usuario
//	@Failure		400	{object}	config.HTTPError
//	@Failure		401	{object}	config.HTTPError
//	@Failure		404	{object}	config.HTTPError
//	@Router			/usuarios/{id} [get]
//	@Security		Bearer
func FindByIdUsuario(c *fiber.Ctx) error {

	id := c.Params("id")

	var usuario model.Usuario

	if checkUsuario(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Usuario não encontrada!"})
	}

	data.DB.Preload("Produto").First(&usuario, id)

	// Limpa o campo de senha antes de retornar o usuário
	usuario.Senha = ""

	return c.Status(fiber.StatusOK).JSON(usuario)
}

//	@Summary		Criar Usuario
//	@Description	Cria um novo Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			usuario	body		model.Usuario	true	"Criar Usuario"
//	@Success		201		{object}	model.Usuario
//	@Failure		400		{object}	config.HTTPError
//	@Failure		404		{object}	config.HTTPError
//	@Router			/usuarios [post]
func CreateUsuario(c *fiber.Ctx) error {

	var usuario *model.Usuario

	if errObjeto := c.BodyParser(&usuario); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	if errValidator := validator.ValidateStruct(usuario); errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errValidator})
	}

	// Verifica se o usuário existe (evitar duplicidade)
	if checkIfUsuarioEmailExists(usuario.Usuario) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Usuário já cadastrado!"})
	}

	// Criptografa a senha
	senhaCriptografada, _ := security.HashPassword(usuario.Senha)
	usuario.Senha = senhaCriptografada

	data.DB.Create(&usuario)

	return c.Status(fiber.StatusCreated).JSON(&usuario)
}

//	@Summary		Atualizar Usuario
//	@Description	Edita um Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			Usuario	body		model.Usuario	true	"Atualizar Usuario"
//	@Success		200		{object}	model.Usuario
//	@Failure		400		{object}	config.HTTPError
//	@Failure		401		{object}	config.HTTPError
//	@Failure		404		{object}	config.HTTPError
//	@Router			/usuarios/atualizar [put]
//	@Security		Bearer
func UpdateUsuario(c *fiber.Ctx) error {

	var usuario *model.Usuario

	if errObjeto := c.BodyParser(&usuario); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	if errValidator := validator.ValidateStruct(usuario); errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errValidator})
	}

	var id = strconv.FormatUint(uint64(usuario.ID), 10)

	// Verifica se o Usuário existe
	if checkUsuario(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Usuário não encontrado!"})
	}

	// Localiza os dados do Usuário
	var buscarUsuario model.Usuario
	data.DB.Where("usuario = ?", usuario.Usuario).First(&buscarUsuario)

	// Verifica se o e-mail já pertence a outro usuário
	if checkIfUsuarioEmailExists(usuario.Usuario) && usuario.ID != buscarUsuario.ID{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Usuário já cadastrado!"})
	}

	// Criptografa a senha
	senhaCriptografada, _ := security.HashPassword(usuario.Senha)
	usuario.Senha = senhaCriptografada

	data.DB.Save(&usuario)

	return c.Status(fiber.StatusOK).JSON(&usuario)
}

// Métodos Auxiliares
func checkUsuario(id string) bool {

	var usuario model.Usuario

	data.DB.First(&usuario, id)

	return usuario.ID == 0

}

func checkIfUsuarioEmailExists(usuarioEmail string) bool {

	var usuario model.Usuario
	data.DB.Where("lower(usuario) = lower(?)", usuarioEmail).Find(&usuario)

	return usuario.Usuario != ""

}
