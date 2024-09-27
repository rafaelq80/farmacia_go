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
	emailService *EmailService // Serviço de envio de e-mail
}

func NewUsuarioService(emailService *EmailService) *UsuarioService {
	return &UsuarioService{
		emailService: emailService,
	}
}

func (usuarioService *UsuarioService) FindAll() []model.Usuario {
	var usuarios []model.Usuario
	data.DB.Preload("Produto").Preload("Role").Omit("Senha").Find(&usuarios)
	return usuarios
}

func (usuarioService *UsuarioService) FindById(id string) (model.Usuario, bool) {
	var usuario model.Usuario
	resposta := data.DB.Preload("Produto").Preload("Role").Omit("Senha").First(&usuario, id)
	return usuario, resposta.RowsAffected > 0
}

func (usuarioService *UsuarioService) FindByUsuario(usuario string) (model.Usuario, bool) {
	var buscaUsuario model.Usuario
	resposta := data.DB.Preload("Produto").Preload("Role").Where("lower(usuario) LIKE lower(?)", "%"+usuario+"%").Find(&buscaUsuario)
	return buscaUsuario, resposta.RowsAffected > 0
}

func (usuarioService *UsuarioService) Create(usuario *model.Usuario) error {

	// Criptografa a senha
	senhaCriptografada, _ := security.HashPassword(usuario.Senha)
	usuario.Senha = senhaCriptografada

	err := data.DB.Create(usuario).Error
	if err != nil {
		return err
	}

	// Enviar E-mail de Confirmação do Cadastro
	subject := "Seja Bem-Vinde ao Projeto Farmácia"
	err = usuarioService.emailService.SendEmail(usuario.Usuario, usuario.Name, subject)
	if err != nil {
		// Exibe o erro no envio, caso aconteça
		fmt.Printf("Failed to send welcome email: %v\n", err)
	}

	return nil
}

func (usuarioService *UsuarioService) Update(usuario *model.Usuario) error {

	// Criptografa a senha
	senhaCriptografada, _ := security.HashPassword(usuario.Senha)
	usuario.Senha = senhaCriptografada

	return data.DB.Save(usuario).Error
}

func (usuarioService *UsuarioService) AutenticarUsuario(usuarioLogin *model.UsuarioLogin) (*model.UsuarioLogin, error) {

	if !usuarioService.EmailExists(usuarioLogin.Usuario) {
		return nil, errors.New("usuário não encontrado")
	}

	// Localiza os dados do Usuário
	usuario, _ := usuarioService.FindByUsuario(usuarioLogin.Usuario)

	// Verifica a Senha
	if !security.CheckPasswordHash(usuarioLogin.Senha, usuario.Senha) {
		return nil, errors.New("senha incorreta")
	}

	// Gera o Token
	token, err := auth.CreateToken(usuarioLogin.Usuario)

	if err != nil {
		return nil, errors.New("erro ao gerar o token")
	}

	log.Printf("Role: %s", usuario.Role.Descricao)

	// Retorna o Objeto usuarioLogin Preenchido
	usuarioLogin.ID = usuario.ID
	usuarioLogin.Nome = usuario.Name
	usuarioLogin.Foto = usuario.Foto
	usuarioLogin.Senha = ""
	usuarioLogin.Role = usuario.Role.Descricao // Role (Não esquecer do Preload da Entidade)
	usuarioLogin.Token = "Bearer " + token

	return usuarioLogin, nil
}

func (usuarioService *UsuarioService) Exists(id string) bool {
	var usuario model.Usuario
	data.DB.First(&usuario, id)
	return usuario.ID != 0
}

func (usuarioService *UsuarioService) EmailExists(usuarioEmail string) bool {
	var usuario model.Usuario
	data.DB.Where("lower(usuario) = lower(?)", usuarioEmail).Find(&usuario)
	return usuario.Usuario != ""
}
