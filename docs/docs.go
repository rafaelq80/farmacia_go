// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Rafael Queiróz",
            "url": "https://github.com/rafaelq80",
            "email": "rafaelproinfo@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categorias": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista todas as Categorias",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categorias"
                ],
                "summary": "Listar Categorias",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Categoria"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Edita uma Categoria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categorias"
                ],
                "summary": "Atualizar Categoria",
                "parameters": [
                    {
                        "description": "Atualizar Categoria",
                        "name": "Categoria",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Categoria"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Categoria"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Cria uma nova Categoria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categorias"
                ],
                "summary": "Criar Categoria",
                "parameters": [
                    {
                        "description": "Criar Categoria",
                        "name": "categoria",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Categoria"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Categoria"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/categorias/grupo/{grupo}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista todas as Categorias por grupo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categorias"
                ],
                "summary": "Listar Categorias por grupo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Grupo do Medicamento (Antibiótico, Antihistamínico, entre outros)",
                        "name": "grupo",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Categoria"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/categorias/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista uma Categoria por id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categorias"
                ],
                "summary": "Listar Categoria por id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id da Categoria",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Categoria"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Apaga uma Categoria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categorias"
                ],
                "summary": "Deletar Categoria",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id da Categoria",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/produtos": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista todos os Produtos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produtos"
                ],
                "summary": "Listar Produtos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Produto"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Edita um Produto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produtos"
                ],
                "summary": "Atualizar Produto",
                "parameters": [
                    {
                        "description": "Atualizar Produto",
                        "name": "Produto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Produto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Produto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Cria um novo Produto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produtos"
                ],
                "summary": "Criar Produto",
                "parameters": [
                    {
                        "description": "Criar Produto",
                        "name": "produto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Produto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Produto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/produtos/nome/{nome}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista todas os Produtos por nome",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produtos"
                ],
                "summary": "Listar Produtos por nome",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome do Produto",
                        "name": "nome",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Produto"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/produtos/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista um Produto por id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produtos"
                ],
                "summary": "Listar Produto por id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id do Produto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Produto"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Apaga um Produto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produtos"
                ],
                "summary": "Deletar Produto",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id do Produto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/usuarios": {
            "post": {
                "description": "Cria um novo Usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Criar Usuario",
                "parameters": [
                    {
                        "description": "Criar Usuario",
                        "name": "usuario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Usuario"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Usuario"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/usuarios/all": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista todos os Usuarios",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Listar Usuarios",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Usuario"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/usuarios/atualizar": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Edita os dados de um Usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Atualizar Usuario",
                "parameters": [
                    {
                        "description": "Atualizar Usuario",
                        "name": "Usuario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Usuario"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Usuario"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/usuarios/logar": {
            "post": {
                "description": "Autentica um Usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Autenticar Usuario",
                "parameters": [
                    {
                        "description": "Autenticar Usuario",
                        "name": "usuario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UsuarioLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UsuarioLogin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        },
        "/usuarios/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Lista um Usuario por id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Listar Usuario por id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id do Usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Usuario"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/config.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.Categoria": {
            "type": "object",
            "required": [
                "grupo"
            ],
            "properties": {
                "grupo": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                },
                "id": {
                    "type": "integer"
                },
                "produto": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Produto"
                    }
                }
            }
        },
        "model.Produto": {
            "type": "object",
            "required": [
                "categoria_id",
                "nome",
                "preco",
                "usuario_id"
            ],
            "properties": {
                "categoria": {
                    "$ref": "#/definitions/model.Categoria"
                },
                "categoria_id": {
                    "type": "integer",
                    "example": 1
                },
                "foto": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                },
                "preco": {
                    "type": "number"
                },
                "usuario": {
                    "$ref": "#/definitions/model.Usuario"
                },
                "usuario_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.Usuario": {
            "type": "object",
            "required": [
                "name",
                "senha",
                "usuario"
            ],
            "properties": {
                "foto": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "produto": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Produto"
                    }
                },
                "senha": {
                    "type": "string"
                },
                "usuario": {
                    "type": "string"
                }
            }
        },
        "model.UsuarioLogin": {
            "type": "object",
            "properties": {
                "foto": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "senha": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "usuario": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "farmacia-go.onrender.com",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "E-commerce - Farmácia",
	Description:      "Projeto E-commerce - Farmácia",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
