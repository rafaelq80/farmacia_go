package test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/server"
	"github.com/rafaelq80/farmacia_go/service"
	"github.com/stretchr/testify/assert"
)

func TestDeveCadastrarProduto(t *testing.T) {

	app := server.SetupServer("teste", true)

	// Injeta o Serviço CategoriaService
	categoriaService := service.NewCategoriaService()

	// Cria o Objeto Categoria
	categoria := model.Categoria{
		ID:    2,
		Grupo: "Antialérgico",
	}

	// Cria a categoria no banco de dados de teste
	err := categoriaService.Create(&categoria)
	assert.NoError(t, err, "Falha ao criar categoria")

	// Cria o Objeto Produto
	produto := model.Produto{
		ID:          1,
		Nome:        "Alektos",
		Preco:       59.90,
		Foto:        "-",
		CategoriaID: categoria.ID,
		UsuarioID:   1,
	}

	// Cria a Requisição HTTP do tipo Post com Autenticação
	resposta, corpoResposta := performRequestWithAuth(t, app, http.MethodPost, "/produtos", produto)

	log.Print(corpoResposta)

	// Verifica se o HTTP Status Code da resposta é o esperado (201)
	assert.Equal(t, fiber.StatusCreated, resposta.StatusCode)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, produto.Nome, corpoResposta["nome"])

}

func TestDeveListarTodasProdutos(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodGet, "/produtos", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveListarProdutoPorId(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodGet, "/produtos/1", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveListarProdutoPorNome(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodGet, "/produtos/nome/anti", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

}

func TestDeveAtualizarProduto(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Injeta o Serviço CategoriaService
	categoriaService := service.NewCategoriaService()

	// Cria o Objeto Categoria
	categoria := model.Categoria{
		ID:    2,
		Grupo: "Antihistamínico",
	}

	// Cria a categoria no banco de dados de teste
	err := categoriaService.Update(&categoria)
	assert.NoError(t, err, "Falha ao atualizar categoria")

	// Atualiza o Objeto Produto
	produto := model.Produto{
		ID:          1,
		Nome:        "Alektos",
		Preco:       69.90,
		Foto:        "-",
		CategoriaID: categoria.ID,
		UsuarioID:   1,
	}

	// Cria a Requisição HTTP do tipo Post com Autenticação
	resposta, corpoResposta := performRequestWithAuth(t, app, http.MethodPut, "/produtos", produto)

	// Verifica se o HTTP Status Code da resposta é o esperado (200)
	assert.Equal(t, fiber.StatusOK, resposta.StatusCode)

	// Verifica no Corpo da Resposta se o usuário foi cadastrado corretamente
	assert.Equal(t, produto.Nome, corpoResposta["nome"])

}

func TestDeveDeletarProdutoPorId(t *testing.T) {

	app := server.SetupServer("teste", false)

	// Cria a Requisição HTTP do tipo Get com Autenticação
	resposta, _ := performRequestWithAuth(t, app, http.MethodDelete, "/produtos/1", nil)

	// Verifica se o HTTP Status Code da resposta é o esperado (204)
	assert.Equal(t, fiber.StatusNoContent, resposta.StatusCode)

}
