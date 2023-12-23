.PHONY: swag-install
swag-install:
	@go install github.com/swaggo/swag/cmd/swag@v1.6.7

.PHONY: build
build:
	@go build -o ./build/app ./src/cmd

.PHONY: swaggo
swaggo:
	@/bin/rm -rf ./docs/swagger
	@`go env GOPATH`/bin/swag init -g ./src/cmd/main.go -o ./docs/swagger --parseInternal	

.PHONY: run
run: swaggo build
	@./build/app