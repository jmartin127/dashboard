# dashboard

This is an example project I put together to demonstrate the utility of some awesome services that exist within the gRPC ecosystem.

# References

gRPC Resources: https://github.com/grpc-ecosystem/awesome-grpc#documentation

Official Website: https://grpc.io/

Helpful Quickstart: https://grpc.io/docs/languages/go/quickstart/

gRPCurl: https://github.com/fullstorydev/grpcurl

# Generating code

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    traffic/traffic.proto

# Making requests

```
$ grpcurl -d '{"origin": "my house", "destination": "Lehi"}' -plaintext localhost:50051 traffic.Traffic/GetTrafficTime
```
