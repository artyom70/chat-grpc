generate-proto:
	protoc --go_opt=paths=source_relative --go_out="." --go-grpc_out=module=chat-task/protos:protos protos/chat.proto

run-server:
	PORT=:8080 go run cmd/server/main.go

run-client-bob:
	PORT=:8080 go run cmd/client/main.go --username=bob

run-client-alisa:
	PORT=:8080 go run cmd/client/main.go --username=alisa

run-client-john:
	PORT=:8080 go run cmd/client/main.go --username=john