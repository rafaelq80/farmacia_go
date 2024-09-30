package setup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	auth "github.com/rafaelq80/farmacia_go/security/service"
	"github.com/rafaelq80/farmacia_go/server"
)

// Variáveis Globais
var (
	testSetup *TestSetup
	setupOnce sync.Once
)

// Dados Compartilhados com todos os testes
type TestSetup struct {
	App   *fiber.App
	Token string
}

// Configuração do Ambiente de testes
func setupTest() (*TestSetup, error) {

	// Inicializa o Servidor de Testes
	app := server.SetupServer("teste", true)

	// Cria o usuário de autenticação dos testes com Token
	usuario := model.Usuario{
		Nome:    "Root User",
		Usuario: "root@email.com",
		Senha:   "rootroot",
		Foto:    "-",
		RoleID:  1,
	}

	usuarioJSON, err := json.Marshal(usuario)
	if err != nil {
		return nil, err
	}

	req := httptest.NewRequest(http.MethodPost, "/usuarios/cadastrar", bytes.NewReader(usuarioJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != fiber.StatusCreated {
		return nil, fmt.Errorf("falha ao criar usuário: status code %d", resp.StatusCode)
	}

	var createdUser model.Usuario
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		return nil, err
	}

	// Gera o Token JWT
	token, err := auth.CreateToken(usuario.Usuario)
	if err != nil {
		return nil, err
	}

	// Compartilha a instância do Servidor de testes e o Token
	return &TestSetup{
		App:   app,
		Token: token,
	}, nil

}

// Método para acessar o Ambiente de testes
func GetTestSetup(t *testing.T) *TestSetup {
	t.Helper()
	setupOnce.Do(func() {
		var err error
		testSetup, err = setupTest()
		if err != nil {
			t.Fatalf("Falha ao configurar o ambiente de teste: %v", err)
		}
	})
	return testSetup
}
