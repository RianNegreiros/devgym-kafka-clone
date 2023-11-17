.PHONY: server client

server:
	cd cmd/server && go run main.go

client:
	cd cmd/client && go run main.go
