gen-grpc:
	rm -f pkg/pb/*.go
	protoc  --proto_path=./api/auth \
		--proto_path=./api/home \
		--go_out=pkg/pb \
		--go_opt=paths=source_relative \
			--go-grpc_out=pkg/pb \
		--go-grpc_opt=paths=source_relative  \
			api/home/*.proto api/auth/*.proto