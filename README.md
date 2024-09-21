# Projeto Farmácia (Em desenvolvimento)

<br />

<div align="center">
    <img src="https://i.imgur.com/JHinCnY.png" title="source: imgur.com" width="10%"/> 
    <img src="https://i.imgur.com/YC6Av6e.png" title="source: imgur.com" /> 
</div>

<br /><br />

## Diagrama de Classes

```mermaid
classDiagram
class Categoria {
  - ID : uint
  - Grupo : string
  - Produto : [] Produto
  + FindAllCategoria()
  + FindByIdCategoria()
  + FindByGrupoCategoria()
  + CreateCategoria()
  + UpdateCategoria()  
  + DeleteCategoria()
}
class Produto {
  - ID : unit
  - Nome : string
  - Preco: float32
  - Foto: string
  - CategoriaID : uint
  - Categoria : *Categoria
  - UsuarioID : uint
  - Usuario : *Usuario
  + FindAllProduto()
  + FindByIdProduto()
  + FindByNomeProduto()
  + CreateProduto()
  + UpdateProduto()  
  + DeleteProduto()
}
class Usuario {
  - ID : uint
  - Name : string
  - Usuario : string
  - Senha : string
  - Foto : string
  - Produto : [] Produto
  + FindAllUsuario()
  + FindByIdUsuario()
  + CreateUsuario()
  + UpdateUsuario()
  + AutenticarUsuario()
}
class UsuarioLogin{
  - Id : uint
  - Nome : string
  - Usuario : string
  - Senha : string
  - Foto : string
  - Token : string
}
Categoria --> Produto
Usuario --> Produto
```

<br /><br />

## Tarefas Concluídas

- [x] Criação do Projeto - Fiber Framework
- [x] Conexão com o Banco de dados MySQL
- [x] Configuração das Rotas
- [x] CRUD de Produto
- [x] CRUD de Categoria
- [x] Relacionamento Produto - Categoria
- [x] CRUD do Usuario
- [x] Relacionamento Produto - Usuario
- [x] Security
- [x] Testes E2E
  - [x] Usuario
  - [x] Categoria
  - [x] Produto
- [x] Swagger
- [x] Deploy no Render
- [x] Refatoramento - Service e Controller
- [x] Envio de E-mails

<br /><br />

# Referências sobre Golang

<br />

<a href="https://go.dev/" target="_blank">Site Oficial - Golang</a>

<a href="https://go.dev/doc/" target="_blank">Documentação Oficial - Golang</a>

<a href="https://pkg.go.dev/" target="_blank">Repositório de pacotes Oficial - Golang</a>

<a href="https://gorm.io/" target="_blank">Biblioteca GORM - Mapeamento Objeto Relacional - Golang</a>

<a href="https://github.com/spf13/viper" target="_blank">Pacote Viper - Gerenciador de configurações da API - Golang</a>

<a href="https://pkg.go.dev/encoding/json" target="_blank">Pacote JSON - Golang</a>

<a href="https://github.com/go-playground/validator" target="_blank">Go Validator V10 - Validação de dados - Golang</a>

<a href="https://github.com/golang-jwt/jwt-docs" target="_blank">Golang JWT - Autenticação com Token JWT - Versão 5.0 - Golang</a>
