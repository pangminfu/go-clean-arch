.PHONY: all test clean build

run: 
	@docker-compose -f deployments/docker-compose.yml up -d app mysql

test:
	@go test ./... || exit 1

build:
	@docker build -f build/package/go/Dockerfile -t go-clean-arch .

run-local:
	@go run cmd/service/main.go

coverage:
	@sh ./test.sh