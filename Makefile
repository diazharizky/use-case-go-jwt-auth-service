.PHONY: run-server run-client mock lint

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go

lint:
	golangci-lint run --timeout 3m

test:
	CONFIG_FILE_PATH=${PWD}/config CONFIG_ENV=test go test -v ./...
