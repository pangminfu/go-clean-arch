.PHONY: all test clean

run: 
	@go run cmd/service/main.go

test:
	@go test ./... || exit 1