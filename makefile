GOROOT=$(shell go env GOROOT)
GOPATH=$(shell go env GOPATH)

PROJECT=fiber-crud-app

SRC=$(wildcard *.go)

TARGET=$(PROJECT)

.PHONY: all build test run clean

all: build

build: $(SRC)
	@go build -o $(TARGET)

test:
	@go test -v ./...

run: build
	@./$(TARGET)

clean:
	@rm -f $(TARGET)
