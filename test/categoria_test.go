package test

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/server"
	"github.com/stretchr/testify/assert"
)

func TestDeveCadastrarCategoria(t *testing.T) {

	app := server.SetupServer("teste", true)

	// Cria o Objeto Categoria
	categoria := model.Categoria{
		ID:    1,
		Grupo: "Antialérgico",
	}

	// Cria a Requisição HTTP do tipo Post com Autenticação
	resposta, corpoResposta := performRequestWithAuth(t, app, http.MethodPost, "/categorias", categoria)

	// Verifica se o HTTP Status Code da resposta é o esperado (201)
	assert.Equal(t, fiber.StatusCreated, resposta.StatusCode)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, categoria.Grupo, corpoResposta["grupo"])

}

func TestDeveListarTodasCategorias(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodGet, "/categorias", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveListarCategoriaPorId(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodGet, "/categorias/1", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveListarCategoriaPorGrupo(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodGet, "/categorias/grupo/anti", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveAtualizarCategoria(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Atualiza o Objeto Categoria
	categoria := model.Categoria{
		ID:    1,
		Grupo: "Antihistamínico",
	}

	// Cria a Requisição HTTP do tipo Post com Autenticação
	resposta, corpoResposta := performRequestWithAuth(t, app, http.MethodPut, "/categorias", categoria)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, categoria.Grupo, corpoResposta["grupo"])

}

func TestDeveDeletarCategoriaPorId(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodDelete, "/categorias/1", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (204)
	assert.Equal(t, fiber.StatusNoContent, resposta.StatusCode)

}
