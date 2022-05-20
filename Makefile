
compile-proto: ## Compile the protobuf contracts to generated Go code
	@protoc --go_out=. --go-grpc_out=. proto/*.proto

run-users-service:
	@go run cmd/users/main.go -http_port=9001 -grpc_port=50051
run-pe-service:
	@go run cmd/physical_environment/main.go -http_port=9003 -grpc_port=50052
run-group-service:
	@go run cmd/groups/main.go -http_port=9004 -grpc_port=50053
up-infra:
	@docker-compose up -d
down-infra:
	@docker-compose down

