package service

import (
	"errors"
	"fmt"

	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
)

type CategoriaService struct{}

func NewCategoriaService() *CategoriaService {
	return &CategoriaService{}
}

func (s *CategoriaService) FindAll() ([]model.Categoria, error) {
	var categorias []model.Categoria
	result := data.DB.Preload("Produto").Find(&categorias)
	return categorias, result.Error
}

func (s *CategoriaService) FindById(id string) (model.Categoria, error) {
	var categoria model.Categoria
	result := data.DB.Preload("Produto").First(&categoria, id)
	if result.RowsAffected == 0 {
		return categoria, errors.New("categoria não encontrada")
	}
	return categoria, result.Error
}

func (s *CategoriaService) FindByGrupo(grupo string) ([]model.Categoria, error) {
	var categorias []model.Categoria
	result := data.DB.Preload("Produto").Where("lower(grupo) LIKE lower(?)", "%"+grupo+"%").Find(&categorias)
	return categorias, result.Error
}

func (s *CategoriaService) Create(categoria *model.Categoria) error {
	if err := data.DB.Create(categoria).Error; err != nil {
		return fmt.Errorf("erro ao criar categoria: %w", err)
	}
	return nil
}

func (s *CategoriaService) Update(categoria *model.Categoria) error {
	if err := data.DB.Save(categoria).Error; err != nil {
		return fmt.Errorf("erro ao atualizar categoria: %w", err)
	}
	return nil
}

func (s *CategoriaService) Delete(id string) error {
	result := data.DB.Delete(&model.Categoria{}, id)
	if result.RowsAffected == 0 {
		return errors.New("categoria não encontrada")
	}
	return result.Error
}

func (s *CategoriaService) Exists(id string) (bool, error) {
	var count int64
	result := data.DB.Model(&model.Categoria{}).Where("id = ?", id).Count(&count)
	return count > 0, result.Error
}
