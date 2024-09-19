package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/service"
	"github.com/rafaelq80/farmacia_go/validator"
)

// Injeção de Dependências - ProdutoService
type ProdutoController struct {
	service *service.ProdutoService
}

// Método Construtor
func NewProdutoController(service *service.ProdutoService) *ProdutoController {
	return &ProdutoController{service: service}
}

//	@Summary		Listar Produtos
//	@Description	Lista todos os Produtos
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Produto
//	@Failure		401	{object}	config.HTTPError
//	@Router			/produtos [get]
//	@Security		Bearer
func (produtoController *ProdutoController) FindAllProduto(context *fiber.Ctx) error {

	produtos := produtoController.service.FindAll()

	return context.Status(fiber.StatusOK).JSON(produtos)

}

//	@Summary		Listar Produto por id
//	@Description	Lista um Produto por id
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Id do Produto"
//	@Success		200	{array}		model.Produto
//	@Failure		400	{object}	config.HTTPError
//	@Failure		401	{object}	config.HTTPError
//	@Failure		404	{object}	config.HTTPError
//	@Router			/produtos/{id} [get]
//	@Security		Bearer
func (produtoController *ProdutoController) FindByIdProduto(context *fiber.Ctx) error {

	id := context.Params("id")

	produto, found := produtoController.service.FindById(id)

	if !found {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrada!"})
	}

	return context.Status(fiber.StatusOK).JSON(produto)

}

//	@Summary		Listar Produtos por nome
//	@Description	Lista todas os Produtos por nome
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			nome	path		string	true	"Nome do Produto"
//	@Success		200		{array}		model.Produto
//	@Failure		400		{object}	config.HTTPError
//	@Failure		401		{object}	config.HTTPError
//	@Router			/produtos/nome/{nome} [get]
//	@Security		Bearer
func (produtoController *ProdutoController) FindByNomeProduto(context *fiber.Ctx) error {

	nome := context.Params("nome")

	produtos := produtoController.service.FindByNome(nome)

	return context.Status(fiber.StatusOK).JSON(produtos)

}

//	@Summary		Criar Produto
//	@Description	Cria um novo Produto
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			produto	body		model.Produto	true	"Criar Produto"
//	@Success		201		{object}	model.Produto
//	@Failure		400		{object}	config.HTTPError
//	@Failure		401		{object}	config.HTTPError
//	@Router			/produtos [post]
//	@Security		Bearer
func (produtoController *ProdutoController) CreateProduto(context *fiber.Ctx) error {

	var produto model.Produto

	if err := context.BodyParser(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	if err := produtoController.service.Create(&produto); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error creating produto"})
	}

	return context.Status(fiber.StatusCreated).JSON(&produto)

}

//	@Summary		Atualizar Produto
//	@Description	Edita um Produto
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			Produto	body		model.Produto	true	"Atualizar Produto"
//	@Success		200		{object}	model.Produto
//	@Failure		400		{object}	config.HTTPError
//	@Failure		401		{object}	config.HTTPError
//	@Failure		404		{object}	config.HTTPError
//	@Router			/produtos [put]
//	@Security		Bearer
func (produtoController *ProdutoController)  UpdateProduto(context *fiber.Ctx) error {

	var produto model.Produto

	if err := context.BodyParser(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	id := strconv.FormatUint(uint64(produto.ID), 10)

	if !produtoController.service.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrada!"})
	}

	if err := produtoController.service.Update(&produto); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error updating produto"})
	}

	return context.Status(fiber.StatusOK).JSON(&produto)

}

//	@Summary		Deletar Produto
//	@Description	Apaga um Produto
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Id do Produto"
//	@Success		204	{string}	string
//	@Failure		400	{object}	config.HTTPError
//	@Failure		401	{object}	config.HTTPError
//	@Failure		404	{object}	config.HTTPError
//	@Router			/produtos/{id} [delete]
//	@Security		Bearer
func (produtoController *ProdutoController) DeleteProduto(context *fiber.Ctx) error {

	id := context.Params("id")

	if !produtoController.service.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrada!"})
	}

	if err := produtoController.service.Delete(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error deleting produto"})
	}

	return context.SendStatus(fiber.StatusNoContent)
}
