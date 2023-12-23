.PHONY: build
build:
	@go build -o ./build/app ./src/cmd

.PHONY: swaggo
swaggo:
	@rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./src/cmd/main.go -o ./docs/swagger --parseInternal	

.PHONY: run
run: swaggo build
	@./build/app