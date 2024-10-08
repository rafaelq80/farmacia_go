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

//	@Summary		Listar Produtos
//	@Description	Lista todos os Produtos
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.Produto
//	@Router			/produtos [get]
//	@Security		Bearer
func (produtoController *ProdutoController) FindAll(context *fiber.Ctx) error {

	produtos, _ := produtoController.produtoService.FindAll()

	return context.Status(fiber.StatusOK).JSON(produtos)

}

//	@Summary		Listar Produto por id
//	@Description	Lista um Produto por id
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Id do Produto"
//	@Success		200	{array}	model.Produto
//	@Router			/produtos/{id} [get]
//	@Security		Bearer
func (produtoController *ProdutoController) FindById(context *fiber.Ctx) error {

	id := context.Params("id")

	produto, err := produtoController.produtoService.FindById(id)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Produto não encontrado!",
		})
	}

	return context.Status(fiber.StatusOK).JSON(produto)

}

//	@Summary		Listar Produtos por nome
//	@Description	Lista todas os Produtos por nome
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			nome	path	string	true	"Nome do Produto"
//	@Success		200		{array}	model.Produto
//	@Router			/produtos/nome/{nome} [get]
//	@Security		Bearer
func (produtoController *ProdutoController) FindByNome(context *fiber.Ctx) error {

	nome := context.Params("nome")

	produtos, err := produtoController.produtoService.FindByNome(nome)

	if err != nil || len(produtos) == 0 {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Nenhum Produto foi encontrado!",
		})
	}

	return context.Status(fiber.StatusOK).JSON(produtos)

}

//	@Summary		Criar Produto
//	@Description	Cria um novo Produto
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			produto	body		model.Produto	true	"Criar Produto"
//	@Success		201		{object}	model.Produto
//	@Router			/produtos [post]
//	@Security		Bearer
func (produtoController *ProdutoController) Create(context *fiber.Ctx) error {

	var produto model.Produto

	if err := context.BodyParser(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err,
		})
	}

	// Validar se a categoria existe
	categoriaId := strconv.FormatUint(uint64(produto.CategoriaID), 10)

	categoriaExiste, _ := produtoController.categoriaService.ExistsById(categoriaId)
	if !categoriaExiste {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Categoria não encontrada!",
		})
	}

	if err := produtoController.produtoService.Create(&produto); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": err.Error(),
		})
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
//	@Router			/produtos [put]
//	@Security		Bearer
func (produtoController *ProdutoController) Update(context *fiber.Ctx) error {

	var produto model.Produto

	if err := context.BodyParser(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(&produto); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": err,
		})
	}

	id := strconv.FormatUint(uint64(produto.ID), 10)
	produtoExist, _ := produtoController.produtoService.ExistsById(id)

	if !produtoExist {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Produto não encontrado!",
		})
	}

	// Validar se a categoria existe
	categoriaId := strconv.FormatUint(uint64(produto.CategoriaID), 10)

	categoriaExiste, _ := produtoController.categoriaService.ExistsById(categoriaId)
	if !categoriaExiste {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Categoria não encontrada!",
		})
	}

	if err := produtoController.produtoService.Update(&produto); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": err.Error(),
		})
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
//	@Router			/produtos/{id} [delete]
//	@Security		Bearer
func (produtoController *ProdutoController) Delete(context *fiber.Ctx) error {

	id := context.Params("id")
	produtoExist, _ := produtoController.produtoService.ExistsById(id)

	if !produtoExist {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "404",
			"message": "Produto não encontrado!",
		})
	}

	if err := produtoController.produtoService.Delete(id); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "500",
			"message": err.Error(),
		})
	}

	return context.SendStatus(fiber.StatusNoContent)
}
