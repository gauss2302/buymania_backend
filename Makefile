build:
	@go build -o bin/buymania cmd/main.go

run: build
	@./bin/buymania

test:
	@go test -v ./...