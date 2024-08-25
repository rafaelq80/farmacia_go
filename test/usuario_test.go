package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	auth "github.com/rafaelq80/farmacia_go/security/service"
	"github.com/rafaelq80/farmacia_go/server"
	"github.com/stretchr/testify/assert"
)

func TestDeveCadastrarUsuario(t *testing.T) {

	// Inicializa o Servidor de Testes
	app := server.SetupServer("teste", true)

	// Cria o Objeto Usuário
	usuario := model.Usuario{
		ID: 	 1,	
		Name:    "Administrador",
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
		Foto:    "-",
	}

	// Converte o Objeto Usuário para o formato JSON
	usuarioJSON, _ := json.Marshal(usuario)

	// Cria a Requisição HTTP do tipo Post
	requisicao := httptest.NewRequest(http.MethodPost, "/usuarios/cadastrar", bytes.NewReader(usuarioJSON))
	
	// Indica no Cabeçalho da Requisição o tipo de dado do Corpo (JSON)
	requisicao.Header.Set("Content-Type", "application/json")

	// Executa a Requisição
	resposta, _ := app.Test(requisicao)

	// Verifica se o HTTP Status Code da resposta é o esperado (201)
	assert.Equal(t, fiber.StatusCreated, resposta.StatusCode)

	// Acessa o Corpo da Requisição
	var corpoResposta interface{}
	json.NewDecoder(resposta.Body).Decode(&corpoResposta)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, usuario.Usuario, corpoResposta.(map[string]interface{})["usuario"])

}

func TestNaoDeveDuplicarUsuario(t *testing.T) {

	app := server.SetupServer("teste", false)

	usuario := model.Usuario{
		ID: 	 1,
		Name:    "Administrador",
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
		Foto:    "-",
	}

	usuarioJSON, _ := json.Marshal(usuario)

	requisicao := httptest.NewRequest(http.MethodPost, "/usuarios/cadastrar", bytes.NewReader(usuarioJSON))
	requisicao.Header.Set("Content-Type", "application/json")

	resposta, _ := app.Test(requisicao)

	// Verifica se o HTTP Status Code da resposta é o esperado (400)
	assert.Equal(t, fiber.StatusBadRequest, resposta.StatusCode)

}

func TestDeveAutenticarUsuario(t *testing.T) {

	app := server.SetupServer("teste", false)

	usuarioLogin := model.UsuarioLogin{
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
	}

	usuarioLoginJSON, _ := json.Marshal(usuarioLogin)

	requisicao := httptest.NewRequest(http.MethodPost, "/usuarios/logar", bytes.NewReader(usuarioLoginJSON))
	requisicao.Header.Set("Content-Type", "application/json")

	resposta, _ := app.Test(requisicao)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveListarTodosUsuarios(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Gera um Token JWT para o Usuário Autenticado
	token,_ := auth.CreateToken("admin@email.com.br")

	requisicao := httptest.NewRequest(http.MethodGet, "/usuarios/all", nil)
	requisicao.Header.Set("Content-Type", "application/json")
	
	// Adiciona o Token JWT no Cabeçalho da Requisição
	requisicao.Header.Set("Authorization", "Bearer " + token)

	resposta, _ := app.Test(requisicao)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveAtualizarUsuario(t *testing.T) {

	app := server.SetupServer("teste", false)

	usuario := model.Usuario{
		ID: 	 1,	
		Name:    "Administrador do Sistema",
		Usuario: "admin@email.com.br",
		Senha:   "admin123",
		Foto:    "-",
	}

	usuarioJSON, _ := json.Marshal(usuario)

	token,_ := auth.CreateToken("admin@email.com.br")

	requisicao := httptest.NewRequest(http.MethodPut, "/usuarios/atualizar", bytes.NewReader(usuarioJSON))
	requisicao.Header.Set("Content-Type", "application/json")
	requisicao.Header.Set("Authorization", "Bearer " + token)

	resposta, _ := app.Test(requisicao)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

	// Acessa o Corpo da Requisição
	var corpoResposta interface{}
	json.NewDecoder(resposta.Body).Decode(&corpoResposta)

	log.Println(corpoResposta)

	// Verifica no Corpo da Resposta se o nome do usuário foi alterado corretamente
	assert.Equal(t, usuario.Name, corpoResposta.(map[string]interface{})["name"])

}

func TestDeveListarUsuarioPorId(t *testing.T) {

	app := server.SetupServer("teste", false)

	token,_ := auth.CreateToken("admin@email.com.br")

	requisicao := httptest.NewRequest(http.MethodGet, "/usuarios/1", nil)
	requisicao.Header.Set("Content-Type", "application/json")
	requisicao.Header.Set("Authorization", "Bearer " + token)

	resposta, _ := app.Test(requisicao)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}