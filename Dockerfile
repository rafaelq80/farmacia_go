# syntax=docker/dockerfile:1

## PASSO 01 - BUILD ##

# Especifica a versão da Linguagem Go
FROM golang:1.23.0-alpine AS build

# Cria a pasta raiz
WORKDIR /app

# Copia os arquivos que contém a lista de dependências
COPY go.mod .
COPY go.sum .

# Instala o Swaggo CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Adiciona o compilador do Golang no Path
ENV PATH=$PATH:/go/bin

# Faz o download de todas as dependências
RUN go mod download

# Copia todos os arquivos para a pasta raiz
COPY . . 

# Verifica se o swagger foi instalado corretamente
RUN which swag

# Libera o acesso para execução de comandos no Terminal
RUN chmod -R 777 .

# Recria a documentação no Swagger
RUN swag fmt 
RUN swag init 

# Compila o projeto (Criar o Build)
RUN go build -o /farmacia_go

## PASSO 02 - DEPLOY ##

# Define a imagem base do Linux para o contêiner. 
FROM alpine:latest

# Define a pasta de trabalho dentro do contêiner.
WORKDIR /

# Copia a pasta /farmacia_go do build para a pasta de trabalho.
COPY --from=build /farmacia_go .

# Indica que o contêiner irá expor a porta 8000.
EXPOSE 8000

# Define o comando principal que será executado para iniciar o projeto.
ENTRYPOINT ["/farmacia_go"]