# BASIC-GO-CRUD-API
A Post managing CRUD API backend app, with OAuth 2.0 auth protocol with JSON-Web-Token. <br/>
Developed with Go language, using Gin web framework &amp; PostgreSQL as database. Also UUID for Post & User ID identifier.


## Tech-stack using :
* Go : programming language
  * Gin : backend web framework to manage routes
  * Gorm : Go ORM library to manage database connection
* PostgreSQL : Using database


## How to run the app
1. In root folder, create `.env` file to holds PORT value variable & PostgreSQL related variables.<br/>
  See `.env.example` for example.
2. Type `$ go get .` command to download required dependencies, that are specified in go.mod file.
3. Type `$ go run main.go` command to run the backend app.


## Run Tests
Run the all tests with `$ go test ./... -v` command, from the root folder.

NOTE : in tests, SQLite database would be created & used instead of PostgreSQL.


## About API endpoints (routes) document and explain
Read endpoints document [here](docs/go-crud-api-endpoint-doc.pdf)
