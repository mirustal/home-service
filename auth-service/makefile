dockerup:
	docker-compose down
	docker-compose up -b
gen-grpc:
	rm -f pkg/pb/*.go
	protoc --proto_path=./api/auth \
    --go_out=./pkg/pb \
    --go_opt=paths=source_relative \
        --go-grpc_out=./pkg/pb \
    --go-grpc_opt=paths=source_relative \
    --experimental_allow_proto3_optional \
        ./api/auth/*.proto

gen-gateway:
	protoc -I . \
          -I ./third-party \
          --go_out=pkg/pb --go_opt=paths=source_relative \
          --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
          --grpc-gateway_out=pkg/pb --grpc-gateway_opt=paths=source_relative \
          --openapiv2_out ./pkg/pb --openapiv2_opt logtostderr=true \
          api/auth/*.proto
