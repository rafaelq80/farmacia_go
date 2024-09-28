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

var (
	testSetup *TestSetup
	setupOnce sync.Once

)

type TestSetup struct {
	App   *fiber.App
	Token string
}

func setupTest() (*TestSetup, error) {

	// Inicializa o Servidor de Testes
	app := server.SetupServer("teste", true)

	// Create a test user
	usuario := model.Usuario{
		Name:    "Root User",
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

	token, err := auth.CreateToken(usuario.Usuario)
	if err != nil {
		return nil, err
	}

	return &TestSetup{
		App:   app,
		Token: token,
	}, nil

}

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
