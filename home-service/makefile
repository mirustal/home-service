gen-grpc:
	protoc -I ./api/home \
		--proto_path=./api/home \
		--go_out=./pkg/pb \
		--go_opt=paths=source_relative \
			--go-grpc_out=./pkg/pb \
		--go-grpc_opt=paths=source_relative \
		--experimental_allow_proto3_optional \
			./api/home/*.proto