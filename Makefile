all: install build test

install:
	go mod download

build:
	go build -o bin/app main.go

run: build
	./bin/app

test:
	go test ./... -v