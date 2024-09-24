# Projeto Farmácia - Rest API
<br />

<div align="center">
    <img src="https://i.imgur.com/JHinCnY.png" title="source: imgur.com" width="15%"/> 
    <img src="https://i.imgur.com/YC6Av6e.png" title="source: imgur.com" /> 
</div>
<br /><br />

Este projeto é um sistema de gerenciamento para farmácias, do tipo MVP, implementado através da Linguagem Go. O sistema visa fornecer uma solução robusta e eficiente para o controle de estoque e gestão de Usuários em estabelecimentos farmacêuticos.

<br />

## Principais Características

1. **Gestão de Estoque**: 
   - Cadastro e atualização de medicamentos
   - Organização por grupo de medicamentos
3. **Gestão de Usuários - Clientes e Fornecedores**: 
   - Cadastro de clientes e fornecedores
   - Envio de e-mail de confirmação do cadastro
4. **Segurança**: 
   - Autenticação e Validação de Token JWT dos usuários

<br />

## Principais Tecnologias Utilizadas

- **Linguagem**: Go
- **Banco de Dados**:
  - MySQL (Desenvolvimento)
  - Sqlite (Testes)
  -  PostgreSQL (Deploy - Render)

- **Framework Web**: Fiber
- **ORM**: GORM

<br />

## Arquitetura

O projeto segue uma arquitetura em camadas, separando as responsabilidades em:

- Controladores (handlers para requisições HTTP)
- Serviços (lógica de negócios)
- Repositórios (interação com o banco de dados)
- Modelos (estruturas de dados)

Esta arquitetura promove a manutenibilidade e escalabilidade do sistema.

<br />

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
