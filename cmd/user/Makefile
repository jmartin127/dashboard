gen:
	cd ../../proto && \
    protoc -I . \
     --go_out=./gen/go/ --go_opt=paths=source_relative \
     --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative \
     --grpc-gateway_out=./gen/go/ --grpc-gateway_opt paths=source_relative \
     --grpc-gateway_opt logtostderr=true \
     --validate_out=lang=go,paths=source_relative:./gen/go \
     jmartin127/user/v1/user.proto
