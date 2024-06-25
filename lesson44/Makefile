.PHONY: protos

protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/translate/translate.proto

run-client:
	go run client/client.go 

run-server:
	go run server/main.go
