package test

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/test/setup"
	"github.com/stretchr/testify/assert"
)

func TestDeveCadastrarUsuario(t *testing.T) {

	// Cria o Objeto Usuário
	usuario := model.Usuario{
		ID:      0,
		Nome:    "Administrador",
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
		Foto:    "-",
		RoleID:  1,
	}

	// Cria a Requisição HTTP do tipo Post com Autenticação
	resposta, corpoResposta := setup.Request(t, http.MethodPost, "/usuarios/cadastrar", usuario)

	// Verifica se o HTTP Status Code da resposta é o esperado (201)
	assert.Equal(t, fiber.StatusCreated, resposta.StatusCode)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, usuario.Nome, corpoResposta["nome"])

}

func TestNaoDeveDuplicarUsuario(t *testing.T) {

	// Cria o Objeto Usuário
	usuario := model.Usuario{
		ID:      0,
		Nome:    "Administrador",
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
		Foto:    "-",
		RoleID:  1,
	}

	// Cria a Requisição HTTP do tipo Post com Autenticação
	resposta, _ := setup.Request(t, http.MethodPost, "/usuarios/cadastrar", usuario)

	// Verifica se o HTTP Status Code da resposta é o esperado (400)
	assert.Equal(t, fiber.StatusBadRequest, resposta.StatusCode)

}

func TestDeveAutenticarUsuario(t *testing.T) {

	usuarioLogin := model.UsuarioLogin{
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
	}

	resposta, _ := setup.Request(t, http.MethodPost, "/usuarios/logar", usuarioLogin)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveListarTodosUsuarios(t *testing.T) {

	resposta, _ := setup.RequestWithAuth(t, http.MethodGet, "/usuarios/all", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveAtualizarUsuario(t *testing.T) {

	usuarioUpdate := model.Usuario{
		ID:      2,
		Nome:    "Administrador do Sistema",
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
		Foto:    "-",
		RoleID:  1,
	}

	resposta, corpoResposta := setup.RequestWithAuth(t, http.MethodPut, "/usuarios/atualizar", usuarioUpdate)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, usuarioUpdate.Nome, corpoResposta["nome"])

}

func TestDeveListarUsuarioPorId(t *testing.T) {

	resposta, _ := setup.RequestWithAuth(t, http.MethodGet, "/usuarios/2", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}
