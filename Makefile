build:
	@go build -o bin/keylang cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go