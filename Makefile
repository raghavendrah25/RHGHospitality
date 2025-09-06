build:
	@go build -o bin/server main.go

run: build
	@./bin/server

test:
	@go test -v ./...