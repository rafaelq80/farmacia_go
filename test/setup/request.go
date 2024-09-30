package setup

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Método para enviar Requisições sem Token
func Request(t *testing.T, method, path string, body ...interface{}) (*http.Response, map[string]interface{}) {
	
	// Indica que a função Request não é uma função de teste
	t.Helper()

	// Obter Ambiente de teste
	setup := GetTestSetup(t)

	// Criar a Requisição
	requisicao := prepareRequest(t, method, path, body...)

	// Enviar a Requisição
	resposta, err := setup.App.Test(requisicao)
	require.NoError(t, err, "Falha ao realizar a requisição")

	// Processar a Resposta
	corpoResposta, err := parseResponse(resposta)
	require.NoError(t, err, "Falha ao analisar a resposta")

	// Retorna a Resposta (HTTP Status) e  Corpo da Resposta
	return resposta, corpoResposta
}

// Método para enviar Requisições com Token
func RequestWithAuth(t *testing.T, method, path string, body ...interface{}) (*http.Response, map[string]interface{}) {

	// Indica que a função RequestWithAuth não é uma função de teste
	t.Helper()

	// Obter Ambiente de teste
	setup := GetTestSetup(t)

	// Criar a Requisição
	requisicao := prepareRequest(t, method, path, body...)

	// Adiciona o Token na Requisição
	requisicao.Header.Set("Authorization", "Bearer "+setup.Token)

	
	// Enviar a Requisição
	resposta, err := setup.App.Test(requisicao)
	require.NoError(t, err, "Falha ao realizar a requisição")

	// Processar a Resposta
	corpoResposta, err := parseResponse(resposta)
	require.NoError(t, err, "Falha ao analisar a resposta")

	// Retorna a Resposta (HTTP Status) e  Corpo da Resposta
	return resposta, corpoResposta

}

// Função para Criar a Requisição
func prepareRequest(t *testing.T, method, path string, body ...interface{}) *http.Request {
	
	// Indica que a função prepareRequest não é uma função de teste
	t.Helper()

	// Cri um ponteiro para um Objeto do tipo http.Request
	var requisicao *http.Request

	// Verifica se algo foi passado no Corpo da Requisição
	if len(body) > 0 {

		// Convert o Corpo da Requisição em JSON
		jsonBody, err := json.Marshal(body[0])
		require.NoError(t, err, "Failed to marshal request body")

		// Cria a Requisicao HTTP
		requisicao = httptest.NewRequest(method, path, bytes.NewReader(jsonBody))

		// Informa que o conteúdo da Requisição é do tipo JSON
		requisicao.Header.Set("Content-Type", "application/json")
	} else {

		// Caso a Requisição não possua Corpo, cria uma nova Requisição sem o Body
		requisicao = httptest.NewRequest(method, path, nil)
	}

	// Retorna a Requisição criada
	return requisicao
}

// Esta função processa respostas HTTP e lida com casos onde o corpo está vazio ou não é um JSON válido.
// Desta forma ela garante que o Corpo da Resposta seja algo que possa ser lido.
func parseResponse(resposta *http.Response) (map[string]interface{}, error) {

	// Se a resposta estiver vazia, a função retorna nulo para ambos os valores de retorno.
	if resposta.ContentLength == 0 {
		return nil, nil
	}

	// Lê todo o conteúdo do corpo da resposta e o armazena no array bodyBytes.
	bodyBytes, err := io.ReadAll(resposta.Body)
	if err != nil {
		return nil, err
	}

	// Garante que o corpo da resposta será fechado quando a função terminar a leitura.
	defer resposta.Body.Close()

	// Recria o Corpo da Resposta
	resposta.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Cria um map para armazenar o corpo da resposta no formato JSON.
	var corpoResposta map[string]interface{}

	// Converte o Corpo da Resposta recriado em JSON
	if err := json.Unmarshal(bodyBytes, &corpoResposta); err != nil {
		
		// Se a conversão falhar, cria um map com a chave "rawBody" contendo o corpo na forma de uma string.
		corpoResposta = map[string]interface{}{"rawBody": string(bodyBytes)}
	}

	// Retornano Corpo da Resposta sem erros
	return corpoResposta, nil
	
}
