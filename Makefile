
server:
	go run cmd/server/main.go -port 8080
test:
	go test -cover -race ./...
