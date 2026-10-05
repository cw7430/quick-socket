set dotenv-load

default:
    just --list

install:
    go mod tidy

dev:
    go run ./cmd/server

build:
    go build -o ./bin/app ./cmd/server

[windows]
clean:
    go clean
    if exist bin rmdir /s /q bin

[unix]
clean:
    go clean
    rm -rf ./bin