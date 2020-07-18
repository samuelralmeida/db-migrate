# Migrate Micro App

Simple Golang software to migrate database using Gorm lib (github.com/jinzhu/gorm).

It uses Postgres database, but the developer can change the db system easily, Gorm supports MySQL, SQLite3 and SQL Server too (https://gorm.io/docs/connecting_to_the_database.html)

## Utility

This initial code can be used directly or adapted in bigger projects.
Like this developer doesn't worry how to migrate his/her database content. She/he can spend more time focusing in the main idea of the project.

## Env file

The software loads database params from `.env` file in root directory using Viper lib (github.com/spf13/viper). Below a `.env` example:

```
HOST=127.0.0.1
PORT=5432
USERNAME=neo
DATABASE=matrix
PASSWORD=morpheus
SSLMODE=disable
```
## How to use

1. Create the entity struct in package model
2. Add the struct memory address (&) to slice Models in `migrate.go`
3. Run `go run migrate.go`
4. Just it