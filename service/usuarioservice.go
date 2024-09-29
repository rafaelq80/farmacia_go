package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
	security "github.com/rafaelq80/farmacia_go/security/bcrypt"
	auth "github.com/rafaelq80/farmacia_go/security/service"
)

type UsuarioService struct {
	emailService *EmailService
}

func NewUsuarioService(emailService *EmailService) *UsuarioService {
	return &UsuarioService{
		emailService: emailService,
	}
}

func (s *UsuarioService) FindAll() ([]model.Usuario, error) {
	var usuarios []model.Usuario
	resposta := data.DB.Preload("Produto").Preload("Role").Omit("Senha").Find(&usuarios)
	return usuarios, resposta.Error
}

func (s *UsuarioService) FindById(id string) (model.Usuario, error) {
	var usuario model.Usuario

	resposta := data.DB.
		Preload("Produto").
		Preload("Role").
		Omit("Senha").
		First(&usuario, id)

	if resposta.RowsAffected == 0 {
		return usuario, errors.New("usuário não encontrado")
	}
	return usuario, resposta.Error
}

func (s *UsuarioService) FindByUsuario(usuario string) (model.Usuario, error) {
	var buscaUsuario model.Usuario

	resposta := data.DB.
		Preload("Produto").
		Preload("Role").
		Where("LOWER(tb_usuarios.usuario) = LOWER(?)", usuario).
		First(&buscaUsuario)

	if resposta.RowsAffected == 0 {
		return buscaUsuario, errors.New("usuário não encontrado")
	}

	return buscaUsuario, resposta.Error
}

func (s *UsuarioService) Create(usuario *model.Usuario) error {
	senhaCriptografada, err := security.HashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("erro ao criptografar senha: %w", err)
	}
	usuario.Senha = senhaCriptografada

	if err := data.DB.Create(usuario).Error; err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}

	subject := "Seja Bem-Vinde ao Projeto Farmácia"
	if err := s.emailService.SendEmail(usuario.Usuario, usuario.Nome, subject); err != nil {
		log.Printf("Falha ao enviar e-mail de boas-vindas: %v\n", err)
	}

	return nil
}

func (s *UsuarioService) Update(usuario *model.Usuario) error {

	senhaCriptografada, err := security.HashPassword(usuario.Senha)

	if err != nil {
		return fmt.Errorf("erro ao criptografar senha: %w", err)
	}

	usuario.Senha = senhaCriptografada

	if err := data.DB.Save(usuario).Error; err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

func (s *UsuarioService) AutenticarUsuario(usuarioLogin *model.UsuarioLogin) (*model.UsuarioLogin, error) {

	usuario, err := s.FindByUsuario(usuarioLogin.Usuario)

	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !security.CheckPasswordHash(usuarioLogin.Senha, usuario.Senha) {
		return nil, errors.New("senha incorreta")
	}

	token, err := auth.CreateToken(usuarioLogin.Usuario)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar o token: %w", err)
	}

	usuarioLogin.ID = usuario.ID
	usuarioLogin.Nome = usuario.Nome
	usuarioLogin.Foto = usuario.Foto
	usuarioLogin.Senha = ""
	usuarioLogin.Token = "Bearer " + token

	return usuarioLogin, nil
}

func (s *UsuarioService) ExistsById(id string) (bool, error) {
	var count int64
	result := data.DB.Model(&model.Usuario{}).Where("id = ?", id).Count(&count)
	return count > 0, result.Error
}

func (s *UsuarioService) EmailExists(usuarioEmail string) bool {
	var count int64
	data.DB.Model(&model.Usuario{}).Where("lower(tb_usuarios.usuario) = lower(?)", usuarioEmail).Count(&count)
	return count > 0
}
