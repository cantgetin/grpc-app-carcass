proto:
	protoc --go_out=. --go-grpc_out=. --proto_path=./third_party --proto_path=./api ./api/api.proto --validate_out="lang=go:./"

run:
	go run cmd/main.go