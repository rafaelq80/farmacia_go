package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/service"
	"github.com/rafaelq80/farmacia_go/validator"
)

// Injeção de Dependências - CategoriaService
type CategoriaController struct {
	service *service.CategoriaService
}

// Método Construtor
func NewCategoriaController(service *service.CategoriaService) *CategoriaController {
	return &CategoriaController{service: service}
}

// @Summary		Listar Categorias
// @Description	Lista todas as Categorias
// @Tags			categorias
// @Accept			json
// @Produce		json
// @Success		200				{array}		model.Categoria
// @Failure		401				{object}	config.HTTPError
// @Router			/categorias 	[get]
// @Security		Bearer
func (categoriaController *CategoriaController) FindAllCategoria(c *fiber.Ctx) error {

	categorias := categoriaController.service.FindAll()

	return c.Status(fiber.StatusOK).JSON(categorias)

}

// @Summary		Listar Categoria por id
// @Description	Lista uma Categoria por id
// @Tags			categorias
// @Accept			json
// @Produce		json
// @Param			id					path		string	true	"Id da Categoria"
// @Success		200					{array}		model.Categoria
// @Failure		400					{object}	config.HTTPError
// @Failure		401					{object}	config.HTTPError
// @Failure		404					{object}	config.HTTPError
// @Router			/categorias/{id} 	[get]
// @Security		Bearer
func (categoriaController *CategoriaController) FindByIdCategoria(context *fiber.Ctx) error {

	id := context.Params("id")

	categoria, found := categoriaController.service.FindById(id)

	if !found {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
	}

	return context.Status(fiber.StatusOK).JSON(categoria)
}

// @Summary		Listar Categorias por grupo
// @Description	Lista todas as Categorias por grupo
// @Tags			categorias
// @Accept			json
// @Produce		json
// @Param			grupo						path		string	true	"Grupo do Medicamento (Antibiótico, Antihistamínico, entre outros)"
// @Success		200							{array}		model.Categoria
// @Failure		400							{object}	config.HTTPError
// @Failure		401							{object}	config.HTTPError
// @Router			/categorias/grupo/{grupo} 	[get]
// @Security		Bearer
func (categoriaController *CategoriaController) FindByGrupoCategoria(context *fiber.Ctx) error {

	grupo := context.Params("grupo")

	categorias := categoriaController.service.FindByGrupo(grupo)

	return context.Status(fiber.StatusOK).JSON(categorias)

}

// @Summary		Criar Categoria
// @Description	Cria uma nova Categoria
// @Tags			categorias
// @Accept			json
// @Produce		json
// @Param			categoria	body		model.Categoria	true	"Criar Categoria"
// @Success		201			{object}	model.Categoria
// @Failure		400			{object}	config.HTTPError
// @Failure		401			{object}	config.HTTPError
// @Router			/categorias [post]
// @Security		Bearer
func (categoriaController *CategoriaController) CreateCategoria(context *fiber.Ctx) error {

	var categoria model.Categoria

	if err := context.BodyParser(&categoria); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&categoria); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	if err := categoriaController.service.Create(&categoria); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error creating categoria"})
	}

	return context.Status(fiber.StatusCreated).JSON(&categoria)

}

// @Summary		Atualizar Categoria
// @Description	Edita uma Categoria
// @Tags			categorias
// @Accept			json
// @Produce		json
// @Param			Categoria	body		model.Categoria	true	"Atualizar Categoria"
// @Success		200			{object}	model.Categoria
// @Failure		400			{object}	config.HTTPError
// @Failure		401			{object}	config.HTTPError
// @Failure		404			{object}	config.HTTPError
// @Router			/categorias [put]
// @Security		Bearer
func (categoriaController *CategoriaController) UpdateCategoria(context *fiber.Ctx) error {

	var categoria model.Categoria

	if err := context.BodyParser(&categoria); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&categoria); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	id := strconv.FormatUint(uint64(categoria.ID), 10)

	if !categoriaController.service.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
	}

	if err := categoriaController.service.Update(&categoria); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error updating categoria"})
	}

	return context.Status(fiber.StatusOK).JSON(&categoria)

}

// @Summary		Deletar Categoria
// @Description	Apaga uma Categoria
// @Tags			categorias
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id da Categoria"
// @Success		204	{string}	string
// @Failure		400	{object}	config.HTTPError
// @Failure		401	{object}	config.HTTPError
// @Failure		404	{object}	config.HTTPError
// @Router			/categorias/{id} [delete]
// @Security		Bearer
func (categoriaController *CategoriaController) DeleteCategoria(context *fiber.Ctx) error {

	id := context.Params("id")

	if !categoriaController.service.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
	}

	if err := categoriaController.service.Delete(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error deleting categoria"})
	}

	return context.SendStatus(fiber.StatusNoContent)

}
