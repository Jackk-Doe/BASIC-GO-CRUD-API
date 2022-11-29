# BASIC-GO-CRUD-API
A basic CRUD API to manange a post data, developed with Go language, using Gin web framework &amp; PostgreSQL as database.


## Tech-stack using :
* Go : programming language
  * Gin : backend web framework to manage routes
  * Gorm : Go ORM library to manage database connection
* PostgreSQL : Using database


## How to run the app
1. In root folder, create `.env` file to holds PORT value variable & PostgreSQL related variables.<br/>
  See `.env.example` for example.
2. Run `$ go get .` command to download required dependencies, that are specified in go.mod file.
3. Run `$ go run main.go` command to run the backend app.

NOTE : check all the created end points in `router/router.go`


## Run Tests
Run the all tests with `$ go test ./... -v` command, from the root folder.

NOTE : in tests, SQLite database would be created & used instead of PostgreSQL.


## About API endpoints (routes) document and explain
Read endpoints document [here](docs/go-crud-api-endpoint-doc.pdf)
