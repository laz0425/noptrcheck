build:
	go build -o noptrcheck ./cmd/main.go
test:
	go test ./...