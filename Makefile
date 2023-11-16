.PHONY: server client

server:
	cd server && go run main.go

client:
	cd client && go run main.go
