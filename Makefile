.PHONY: generate
generate:
	protoc api/v1/*.proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=.

.PHONY: server
server:
	go run cmd/server/server.go


.PHONY: client
client:
	go run cmd/server/client.go