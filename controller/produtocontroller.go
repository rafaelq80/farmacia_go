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
	produtoService   *service.ProdutoService
	categoriaService *service.CategoriaService
}

// Método Construtor
func NewProdutoController(produtoService *service.ProdutoService, categoriaService *service.CategoriaService) *ProdutoController {
	return &ProdutoController{
		produtoService:   produtoService,
		categoriaService: categoriaService,
	}
}

// @Summary		Listar Produtos
// @Description	Lista todos os Produtos
// @Tags			produtos
// @Accept			json
// @Produce		json
// @Success		200	{array}		model.Produto
// @Failure		401	{object}	config.HTTPError
// @Router			/produtos [get]
// @Security		Bearer
func (produtoController *ProdutoController) FindAllProduto(context *fiber.Ctx) error {

	produtos := produtoController.produtoService.FindAll()

	return context.Status(fiber.StatusOK).JSON(produtos)

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
func (produtoController *ProdutoController) FindByIdProduto(context *fiber.Ctx) error {

	id := context.Params("id")

	produto, found := produtoController.produtoService.FindById(id)

	if !found {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrada!"})
	}

	return context.Status(fiber.StatusOK).JSON(produto)

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
func (produtoController *ProdutoController) FindByNomeProduto(context *fiber.Ctx) error {

	nome := context.Params("nome")

	produtos := produtoController.produtoService.FindByNome(nome)

	return context.Status(fiber.StatusOK).JSON(produtos)

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
func (produtoController *ProdutoController) CreateProduto(context *fiber.Ctx) error {

	var produto model.Produto

	if err := context.BodyParser(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	// Validar se a categoria existe
	categoriaId := strconv.FormatUint(uint64(produto.CategoriaID), 10)

	categoriaExiste := produtoController.categoriaService.Exists(categoriaId)
	if !categoriaExiste {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Categoria não existe"})
	}


	if err := produtoController.produtoService.Create(&produto); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Erro ao criar o produto"})
	}

	return context.Status(fiber.StatusCreated).JSON(&produto)

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
func (produtoController *ProdutoController) UpdateProduto(context *fiber.Ctx) error {

	var produto model.Produto

	if err := context.BodyParser(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err.Error()})
	}

	if err := validator.ValidateStruct(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": err})
	}

	id := strconv.FormatUint(uint64(produto.ID), 10)

	if !produtoController.produtoService.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrada!"})
	}

	if err := produtoController.produtoService.Update(&produto); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error updating produto"})
	}

	return context.Status(fiber.StatusOK).JSON(&produto)

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
func (produtoController *ProdutoController) DeleteProduto(context *fiber.Ctx) error {

	id := context.Params("id")

	if !produtoController.produtoService.Exists(id) {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Produto não encontrada!"})
	}

	if err := produtoController.produtoService.Delete(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "500", "message": "Error deleting produto"})
	}

	return context.SendStatus(fiber.StatusNoContent)
}
