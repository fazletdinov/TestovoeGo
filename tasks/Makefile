.PHONY: proto
proto:
	protoc -I ./proto \
  --go_out=./protogen --go_opt=paths=source_relative \
  --go-grpc_out=./protogen --go-grpc_opt=paths=source_relative \
  ./proto/tasks/tasks.proto

.PHONY: run
run:
	docker compose -f docker-compose.yaml up --build