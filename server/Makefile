run-mcdonald:
	@go run services/mcdonald/main.go services/mcdonald/grpc.go

run-orders:
	@go run services/orders/main.go

run-kitchen:
	@go run services/kitchen/main.go 

gen:
	@protoc \
		--proto_path=proto "proto/mcdonald.proto" \
		--go_out=services/common/genproto/mcdonald --go_opt=paths=source_relative \
  	--go-grpc_out=services/common/genproto/mcdonald --go-grpc_opt=paths=source_relative