package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/validator"
)

//	@Summary		Listar Categorias
//	@Description	Lista todas as Categorias
//	@Tags			categorias
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Categoria
//	@Failure		401	{object}	config.HTTPError
//	@Router			/categorias [get]
//	@Security		Bearer
func FindAllCategoria(c *fiber.Ctx) error {

	var categorias []model.Categoria

	data.DB.Preload("Produto").Find(&categorias)

	return c.Status(fiber.StatusOK).JSON(categorias)
}

//	@Summary		Listar Categoria por id
//	@Description	Lista uma Categoria por id
//	@Tags			categorias
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Id da Categoria"
//	@Success		200	{array}		model.Categoria
//	@Failure		400	{object}	config.HTTPError
//	@Failure		401	{object}	config.HTTPError
//	@Failure		404	{object}	config.HTTPError
//	@Router			/categorias/{id} [get]
//	@Security		Bearer
func FindByIdCategoria(c *fiber.Ctx) error {

	id := c.Params("id")

	var categoria model.Categoria

	if checkCategoria(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
	}

	data.DB.Preload("Produto").First(&categoria, id)

	return c.Status(fiber.StatusOK).JSON(categoria)
}

//	@Summary		Listar Categorias por grupo
//	@Description	Lista todas as Categorias por grupo
//	@Tags			categorias
//	@Accept			json
//	@Produce		json
//	@Param			grupo	path		string	true	"Grupo do Medicamento (Antibiótico, Antihistamínico, entre outros)"
//	@Success		200		{array}		model.Categoria
//	@Failure		400		{object}	config.HTTPError
//	@Failure		401		{object}	config.HTTPError
//	@Router			/categorias/grupo/{grupo} [get]
//	@Security		Bearer
func FindByGrupoCategoria(c *fiber.Ctx) error {

	grupo := c.Params("grupo")

	var categorias []model.Categoria

	data.DB.Preload("Produto").Where("lower(grupo) LIKE lower(?)", "%"+grupo+"%").Find(&categorias)

	return c.Status(fiber.StatusOK).JSON(categorias)
}

//	@Summary		Criar Categoria
//	@Description	Cria uma nova Categoria
//	@Tags			categorias
//	@Accept			json
//	@Produce		json
//	@Param			categoria	body		model.Categoria	true	"Criar Categoria"
//	@Success		201			{object}	model.Categoria
//	@Failure		400			{object}	config.HTTPError
//	@Failure		401			{object}	config.HTTPError
//	@Router			/categorias [post]
//	@Security		Bearer
func CreateCategoria(c *fiber.Ctx) error {

	var categoria *model.Categoria

	if errObjeto := c.BodyParser(&categoria); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	if errValidator := validator.ValidateStruct(categoria); errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errValidator})
	}

	data.DB.Create(&categoria)

	return c.Status(fiber.StatusCreated).JSON(&categoria)
}

//	@Summary		Atualizar Categoria
//	@Description	Edita uma Categoria
//	@Tags			categorias
//	@Accept			json
//	@Produce		json
//	@Param			Categoria	body		model.Categoria	true	"Atualizar Categoria"
//	@Success		200			{object}	model.Categoria
//	@Failure		400			{object}	config.HTTPError
//	@Failure		401			{object}	config.HTTPError
//	@Failure		404			{object}	config.HTTPError
//	@Router			/categorias [put]
//	@Security		Bearer
func UpdateCategoria(c *fiber.Ctx) error {

	var categoria *model.Categoria

	if errObjeto := c.BodyParser(&categoria); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	if errValidator := validator.ValidateStruct(categoria); errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errValidator})
	}

	var id = strconv.FormatUint(uint64(categoria.ID), 10)

	if checkCategoria(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
	}

	data.DB.Save(&categoria)

	return c.Status(fiber.StatusOK).JSON(&categoria)
}

//	@Summary		Deletar Categoria
//	@Description	Apaga uma Categoria
//	@Tags			categorias
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Id da Categoria"
//	@Success		204	{string}	string	
//	@Failure		400	{object}	config.HTTPError
//	@Failure		401	{object}	config.HTTPError
//	@Failure		404	{object}	config.HTTPError
//	@Router			/categorias/{id} [delete]
//	@Security		Bearer
func DeleteCategoria(c *fiber.Ctx) error {

	id := c.Params("id")

	var categoria model.Categoria

	if checkCategoria(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Categoria não encontrada!"})
	}

	data.DB.Delete(&categoria, id)

	return c.SendStatus(fiber.StatusNoContent)
}

// Método Auxiliar
func checkCategoria(id string) bool {

	var categoria model.Categoria

	data.DB.First(&categoria, id)

	return categoria.ID == 0

}
