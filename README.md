[![Golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org)
[![MongoDB](https://www.vectorlogo.zone/logos/mongodb/mongodb-ar21.svg)](https://www.mongodb.com/)  
[![Build Status](https://img.shields.io/travis/archit-p/go-microservice)](https://travis-ci.org/archit-p/go-microservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/archit-p/go-microservice)](https://goreportcard.com/report/github.com/archit-p/go-microservice)  
Go Microservice built using best practices, ideal for use as a starter template. It features an extensible model - Sample - with support for CRUD operations for MongoDB.

## References
The idea for this project came up after reading multiple blogs and guides on best practices while writing Go code. I'm sharing links to these below.
### Project Structure
1. [How to Write Go Code](https://golang.org/doc/code.html)
2. [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
3. [Organizing Database Access](https://www.alexedwards.net/blog/organising-database-access)
### Documentation
1. [Documenting APIs with Swagger](https://swagger.io/resources/articles/documenting-apis-with-swagger/)
2. [Documenting Go Code](https://blog.golang.org/godoc)
### Containerization
1. [Docker: Multi-Stage Builds](https://docs.docker.com/develop/develop-images/multistage-build/)
2. [Why You Should Use Microservices and Containers](https://developer.ibm.com/technologies/microservices/articles/why-should-we-use-microservices-and-containers/)
### Testing
1. [Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c)
2. [Testing Web-Apps in Go](https://markjberger.com/testing-web-apps-in-golang/)
## Project Layout
```
cmd             (contains code for our apps)
|-+ web         (server router and controllers)
pkg             (contains reusable code)
|-+ dto         (data-transfer objects)
|-+ models      (database models)
    |-+ mongodb (models for mongo-db)
|-+ docs        (swagger documentation)
```
## Running
The project ships with a Makefile to build and run the service.
```sh
> make help
Run make <target> where target is
	help: print out this message
	build: build the executables
	run: start a clean build, and run executable
	test: run go tests
	docs: build documentation
	clean: clean executables and docs
> make run
```

## Docs
Once the service is running, accompanying Swagger docs can be found at `http://localhost:8080/swagger/index.html`.

## Contributing
Feel free to fork the project for use as a base template for your next microservice or REST API!
