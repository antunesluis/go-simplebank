# Project Description

This repository contains the application developed during the [Backend Master Class](https://bit.ly/backendmaster).

The project is a simple banking system that allows creating and managing bank accounts, recording all balance changes in each account, and performing transfers between accounts.

## Libraries Used

- [golang-migrate](https://github.com/golang-migrate/migrate): Used for database migratione.
- [SQLC](https://github.com/sqlc-dev/sqlc): Used to generate CRUD code.
- [Testify](https://github.com/stretchr/testify): Testify para escrever testes unitários.
- [Gin Framework](https://github.com/gin-gonic/gin): Used for implementing the REST API.
- [Golang Viper](https://github.com/spf13/viper): Used for configuring files and environment variables.
- [GoMock](https://github.com/uber-go/mock): Used for mocking the database.
- [JWT](https://github.com/golang-jwt/jwt) e [a biblioteca PASETO](https://github.com/o1egl/paseto): Used for creating login tokens.

## Setup local development

### Install tools

- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

  ```bash
  curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
  ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

  ```bash
  sudo snap install sqlc
  ```

- [Gomock](https://github.com/golang/mock)

  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```

### Setup infrastructure

- Start postgres container:

  ```bash
  make postgres
  ```

- Create simple_bank database:

  ```bash
  make createdb
  ```

- Run db migration up all versions:

  ```bash
  make migrateup
  ```

- Run db migration up 1 version:

  ```bash
  make migrateup1
  ```

- Run db migration down all versions:

  ```bash
  make migratedown
  ```

- Run db migration down 1 version:

  ```bash
  make migratedown1
  ```

### How to generate code

- Generate SQL CRUD with sqlc:

  ```bash
  make sqlc
  ```
