package service

import (
	"errors"
	"fmt"

	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) FindAll() ([]model.Role, error) {
	var roles []model.Role
	result := data.DB.Preload("Usuario").Find(&roles)
	return roles, result.Error
}

func (s *RoleService) FindById(id string) (model.Role, error) {
	var role model.Role
	result := data.DB.Preload("Usuario").First(&role, id)
	if result.RowsAffected == 0 {
		return role, errors.New("role não encontrada")
	}
	return role, result.Error
}

func (s *RoleService) Create(role *model.Role) error {
	if err := data.DB.Create(role).Error; err != nil {
		return fmt.Errorf("erro ao criar role: %w", err)
	}
	return nil
}

func (s *RoleService) Update(role *model.Role) error {
	if err := data.DB.Save(role).Error; err != nil {
		return fmt.Errorf("erro ao atualizar role: %w", err)
	}
	return nil
}

func (s *RoleService) ExistsById(id string) (bool, error) {
	var count int64
	result := data.DB.Model(&model.Role{}).Where("id = ?", id).Count(&count)
	return count > 0, result.Error
}

func (s *RoleService) RoleExists(role string) bool {
	var count int64
	data.DB.Model(&model.Role{}).Where("lower(tb_roles.role) = lower(?)", role).Count(&count)
	return count > 0
}