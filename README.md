# BASIC-GO-CRUD-API
A basic CRUD API to manange a post data, developed with Go language, using Gin web framework &amp; PostgreSQL as database.

## Tech-stack using :
* Go : programming language
  * Gin : backend web framework to manage routes
  * Gorm : Go ORM library to manage database connection
  * Godotenv : load .env file
  * CompileDaemon : debug tool (Optional)
* PostgreSQL : Using database

## How to run the app
1. In root folder, create `.env` file to holds PORT value variable & PostgreSQL related variables.<br/>
  See `.env.example` for example.
2. Run `$ go mod tidy` command to download necessary modules, that are specified in go.mod file.
3. Run `$ go migrate/migrate.go` command to (initialize) create table on the database.
4. Run `$ go run main.go` command to run the backend app.

NOTE : check all the created end points in `main.go`
