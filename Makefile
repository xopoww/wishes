generate:
	go generate ./...
	go mod tidy

build:
	go build -race -o build/wishes-server ./cmd/wishes-server

test:
	go test ./...