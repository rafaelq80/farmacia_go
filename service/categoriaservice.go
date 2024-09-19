package service

import (
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
)

type CategoriaService struct{}

func NewCategoriaService() *CategoriaService {
	return &CategoriaService{}
}

func (categoriaService *CategoriaService) FindAll() []model.Categoria {
	var categorias []model.Categoria
	data.DB.Preload("Produto").Find(&categorias)
	return categorias
}

func (categoriaService *CategoriaService) FindById(id string) (model.Categoria, bool) {
	var categoria model.Categoria
	resposta := data.DB.Preload("Produto").First(&categoria, id)
	return categoria, resposta.RowsAffected > 0
}

func (categoriaService *CategoriaService) FindByGrupo(grupo string) []model.Categoria {
	var categorias []model.Categoria
	data.DB.Preload("Produto").Where("lower(grupo) LIKE lower(?)", "%"+grupo+"%").Find(&categorias)
	return categorias
}

func (categoriaService *CategoriaService) Create(categoria *model.Categoria) error {
	return data.DB.Create(categoria).Error
}

func (categoriaService *CategoriaService) Update(categoria *model.Categoria) error {
	return data.DB.Save(categoria).Error
}

func (categoriaService *CategoriaService) Delete(id string) error {
	return data.DB.Delete(&model.Categoria{}, id).Error
}

func (categoriaService *CategoriaService) Exists(id string) bool {
	var categoria model.Categoria
	data.DB.First(&categoria, id)
	return categoria.ID != 0
}