run:
	go run ./cmd/main.go
dev:
	air
lint:
	golangci-lint run ./...
fmt:
	go fmt ./...
vet:
	go vet ./...