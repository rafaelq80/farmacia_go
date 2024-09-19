package service

import (
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
)

type ProdutoService struct{}

func NewProdutoService() *ProdutoService {
	return &ProdutoService{}
}

func (produtoService *ProdutoService) FindAll() []model.Produto {
	var produtos []model.Produto
	data.DB.Joins("Categoria").Joins("Usuario").Find(&produtos)
	return produtos
}

func (produtoService *ProdutoService) FindById(id string) (model.Produto, bool) {
	var produto model.Produto
	resposta := data.DB.Joins("Categoria").Joins("Usuario").First(&produto, id)
	return produto, resposta.RowsAffected > 0
}

func (produtoService *ProdutoService) FindByNome(nome string) []model.Produto {
	var produtos []model.Produto
	data.DB.Joins("Categoria").Joins("Usuario").Where("lower(nome) LIKE lower(?)", "%"+nome+"%").Find(&produtos)
	return produtos
}

func (produtoService *ProdutoService) Create(produto *model.Produto) error {
	return data.DB.Create(produto).Error
}

func (produtoService *ProdutoService) Update(produto *model.Produto) error {
	return data.DB.Save(produto).Error
}

func (produtoService *ProdutoService) Delete(id string) error {
	return data.DB.Delete(&model.Produto{}, id).Error
}

func (produtoService *ProdutoService) Exists(id string) bool {
	var produto model.Produto
	data.DB.First(&produto, id)
	return produto.ID != 0
}