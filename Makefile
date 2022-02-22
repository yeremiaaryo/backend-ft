test:
	go test ./... -cover

build:
	go build -o bin/main main.go

all: test build
