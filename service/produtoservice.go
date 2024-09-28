package service

import (
	"errors"
	"fmt"

	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
)

type ProdutoService struct{}

func NewProdutoService() *ProdutoService {
	return &ProdutoService{}
}

func (s *ProdutoService) FindAll() ([]model.Produto, error) {
	var produtos []model.Produto
	result := data.DB.Joins("Categoria").Joins("Usuario").Find(&produtos)
	return produtos, result.Error
}

func (s *ProdutoService) FindById(id string) (model.Produto, error) {
	var produto model.Produto
	result := data.DB.Joins("Categoria").Joins("Usuario").First(&produto, id)
	if result.RowsAffected == 0 {
		return produto, errors.New("produto não encontrado")
	}
	return produto, result.Error
}

func (s *ProdutoService) FindByNome(nome string) ([]model.Produto, error) {
	var produtos []model.Produto
	result := data.DB.Joins("Categoria").Joins("Usuario").Where("lower(nome) LIKE lower(?)", "%"+nome+"%").Find(&produtos)
	return produtos, result.Error
}

func (s *ProdutoService) Create(produto *model.Produto) error {
	if err := data.DB.Create(produto).Error; err != nil {
		return fmt.Errorf("erro ao criar produto: %w", err)
	}
	return nil
}

func (s *ProdutoService) Update(produto *model.Produto) error {
	if err := data.DB.Save(produto).Error; err != nil {
		return fmt.Errorf("erro ao atualizar produto: %w", err)
	}
	return nil
}

func (s *ProdutoService) Delete(id string) error {
	result := data.DB.Delete(&model.Produto{}, id)
	if result.RowsAffected == 0 {
		return errors.New("produto não encontrado")
	}
	return result.Error
}

func (s *ProdutoService) ExistsById(id string) (bool, error) {
	var count int64
	result := data.DB.Model(&model.Produto{}).Where("id = ?", id).Count(&count)
	return count > 0, result.Error
}
