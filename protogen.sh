protoc --proto_path=./proto --go_out=./pb/customer --go_opt=paths=source_relative \
    --go-grpc_out=./pb/customer --go-grpc_opt=paths=source_relative \
    ./proto/*.proto