gen-gateway:
	rm -f pkg/pb/*.go
	protoc -I . \
          -I ./third-party \
		  -I ./api/auth \
		  -I ./api/home \
          --go_out=pkg/pb --go_opt=paths=source_relative \
          --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
          --grpc-gateway_out=pkg/pb --grpc-gateway_opt=paths=source_relative \
          --openapiv2_out ./pkg/pb --openapiv2_opt logtostderr=true \
          api/auth/*.proto api/home/*.proto
