package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/validator"
)

// @Summary		Listar Produtos
// @Description	Lista todos os Produtos
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Success		200	{array}		model.Produto
// @Failure		401	{object}	config.HTTPError
// @Router			/produtos [get]
// @Security		Bearer
func FindAllProduto(c *fiber.Ctx) error {

	var produtos []model.Produto

	data.DB.Joins("Categoria").Joins("Usuario").Find(&produtos)

	return c.Status(fiber.StatusOK).JSON(produtos)
}

// @Summary		Listar Produto por id
// @Description	Lista um Produto por id
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id do Produto"
// @Success		200	{array}		model.Produto
// @Failure		400	{object}	config.HTTPError
// @Failure		401	{object}	config.HTTPError
// @Failure		404	{object}	config.HTTPError
// @Router			/produtos/{id} [get]
// @Security		Bearer
func FindByIdProduto(c *fiber.Ctx) error {

	id := c.Params("id")

	var produto model.Produto

	if checkProduto(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrado!"})
	}

	data.DB.Joins("Categoria").Joins("Usuario").First(&produto, id)

	return c.Status(fiber.StatusOK).JSON(produto)
}

// @Summary		Listar Produtos por nome
// @Description	Lista todas os Produtos por nome
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Param			nome	path		string	true	"Nome do Produto"
// @Success		200		{array}		model.Produto
// @Failure		400		{object}	config.HTTPError
// @Failure		401		{object}	config.HTTPError
// @Router			/produtos/nome/{nome} [get]
// @Security		Bearer
func FindByNomeProduto(c *fiber.Ctx) error {

	nome := c.Params("nome")

	var produtos []model.Produto

	data.DB.Joins("Categoria").Joins("Usuario").Where("lower(nome) LIKE lower(?)", "%"+nome+"%").Find(&produtos)

	return c.Status(fiber.StatusOK).JSON(produtos)
}

// @Summary		Criar Produto
// @Description	Cria um novo Produto
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Param			produto	body		model.Produto	true	"Criar Produto"
// @Success		201		{object}	model.Produto
// @Failure		400		{object}	config.HTTPError
// @Failure		401		{object}	config.HTTPError
// @Router			/produtos [post]
// @Security		Bearer
func CreateProduto(c *fiber.Ctx) error {

	var produto *model.Produto

	if errObjeto := c.BodyParser(&produto); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	if errValidator := validator.ValidateStruct(produto); errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errValidator})
	}

	if errDatabase := data.DB.Create(&produto).Error; errDatabase != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errDatabase})
	}

	return c.Status(fiber.StatusCreated).JSON(&produto)
}

// @Summary		Atualizar Produto
// @Description	Edita um Produto
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Param			Produto	body		model.Produto	true	"Atualizar Produto"
// @Success		200		{object}	model.Produto
// @Failure		400		{object}	config.HTTPError
// @Failure		401		{object}	config.HTTPError
// @Failure		404		{object}	config.HTTPError
// @Router			/produtos [put]
// @Security		Bearer
func UpdateProduto(c *fiber.Ctx) error {

	var produto *model.Produto

	if errObjeto := c.BodyParser(&produto); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	if errValidator := validator.ValidateStruct(produto); errValidator != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errValidator})
	}

	var id = strconv.FormatUint(uint64(produto.ID), 10)

	if checkProduto(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrado!"})
	}

	if errDatabase := data.DB.Save(&produto).Error; errDatabase != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errDatabase})
	}

	return c.Status(fiber.StatusOK).JSON(&produto)
}

// @Summary		Deletar Produto
// @Description	Apaga um Produto
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Id do Produto"
// @Success		204	{string}	string
// @Failure		400	{object}	config.HTTPError
// @Failure		401	{object}	config.HTTPError
// @Failure		404	{object}	config.HTTPError
// @Router			/produtos/{id} [delete]
// @Security		Bearer
func DeleteProduto(c *fiber.Ctx) error {

	id := c.Params("id")

	var produto model.Produto

	if checkProduto(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrado!"})
	}

	data.DB.Delete(&produto, id)

	return c.SendStatus(fiber.StatusNoContent)
}

// Método Auxiliar
func checkProduto(id string) bool {

	var produto model.Produto

	data.DB.First(&produto, id)

	return produto.ID == 0

}
