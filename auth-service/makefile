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


MOCKS_DESTINATION=mocks
.PHONY: mocks

mocks: pkg/internal/modules/auth/auth.go 
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done