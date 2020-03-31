
#!make

build:
	go build -o bin/parking_lot cmd/main.go

test:
	go test ./...