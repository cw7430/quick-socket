set dotenv-load
set shell := ["sh", "-cu"]
set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

default:
    just --list

install:
    go mod tidy

dev:
    go run ./cmd/server

build:
    go build -o ./bin/app ./cmd/server

[windows]
doc:
     & "$(go env GOPATH)\bin\swag.exe" init -g cmd/server/main.go

[unix]
doc:
    "$(go env GOPATH)/bin/swag" init -g cmd/server/main.go

[windows]
clean:
    go clean
    if exist bin rmdir /s /q bin

[unix]
clean:
    go clean
    rm -rf ./bin