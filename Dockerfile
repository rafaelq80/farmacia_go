# syntax=docker/dockerfile:1

## PASSO 01 - BUILD ##

# Especifica a versão do Go

FROM golang:1.23.0-alpine AS build

# Cria a pasta raiz

WORKDIR /app

# Copia os arquivos que contém a lista de dependências

COPY go.mod .
COPY go.sum .

# Faz o download de todas as dependências

RUN go mod download

# Copia todos os arquivos para a pasta raiz

COPY . . 

# Compila o projeto

RUN go build -o /farmacia_go

## PASSO 02 - DEPLOY ##

FROM alpine:latest

WORKDIR /

COPY --from=build /farmacia_go .

EXPOSE 8000

ENTRYPOINT ["/farmacia_go"]