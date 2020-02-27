.PHONY: all test clean build

run: 
	@go run cmd/service/main.go

test:
	@go test ./... || exit 1

build:
	@docker build -f build/package/go/Dockerfile -t go-clean-arch .