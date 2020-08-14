################## Makefile ###################
GOCMD=go
TARGET=main
SOURCE=./cmd/web/*
DOCS=./docs/docs_init.go
SWAGGERCMD=swag

.PHONY: help dep build test docs clean

help:
	@echo "Run make <target> where target is"
	@echo "	help: print out this message"
	@echo "	build: build the executables"
	@echo "	run: start a clean build, and run executable"
	@echo "	test: run go tests"
	@echo "	docs: build documentation"
	@echo "	clean: clean executables and docs"

clean:
	rm -f ${TARGET} 2>&1 1>/dev/null

docs:
	go get -u github.com/swaggo/swag/cmd/swag
	${SWAGGERCMD} init -g ${DOCS}

build: clean docs
	${GOCMD} build -o ${TARGET} ${SOURCE}

test: dep
	${GOCMD} test ./... -cover -race -coverprofile=cov.out

run: clean build
	./${TARGET}
