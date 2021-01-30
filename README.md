# dashboard

This is an example project I put together to demonstrate the utility of some awesome services that exist within the gRPC ecosystem.

# To Run Locally

`make all`

# Swagger

http://localhost:8081/swaggerui/#/

# Example gRPC requests

List out gRPC functions supported by a service:
```
$ grpcurl -plaintext localhost:50051 list jmartin127.traffic.v1.Traffic
jmartin127.traffic.v1.Traffic.GetTravelTime
```

Retrieve Travel Time:
```
$ grpcurl -d '{"originAddress": "my house", "destinationAddress": "Lehi"}' -plaintext localhost:50051 jmartin127.traffic.v1.Traffic/GetTravelTime
{
  "travelTime": "660s"
}
```

Retrieve Current Weather:
```
$ grpcurl -d '{"address": "lehi"}' -plaintext localhost:50052 jmartin127.weather.v1.Weather/GetCurrentWeather
{
  "tempFahrenheit": 39,
  "precipitationPct": 1,
  "humidityPct": 38,
  "windMPH": 5
}
```

# References

gRPC Resources: https://github.com/grpc-ecosystem/awesome-grpc#documentation

Mapping of gRPC errors to HTTP status codes: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go

Official Website: https://grpc.io/

Helpful Quickstart: https://grpc.io/docs/languages/go/quickstart/

gRPCurl: https://github.com/fullstorydev/grpcurl

grpc-json-proxy: https://github.com/jnewmano/grpc-json-proxy

grpc-gateway: https://github.com/grpc-ecosystem/grpc-gateway

grpc google.api.http annotations: https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46

all the crazy annotations: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/internal/proto/examplepb/a_bit_of_everything.proto

Buff Style Guide: https://docs.buf.build/style-guide/#files-and-packages

Swagger UI: https://www.ribice.ba/serving-swaggerui-golang/

Swagger UI: https://github.com/swagger-api/swagger-ui

Validation (protoc-gen-validate): https://github.com/envoyproxy/protoc-gen-validate
# TODO
1. Add an Aggergation Dashboard service
1. Add a User service
1. Fix the title fo the swagger page
1. Finish the presentation
1. Actually hook this up to Google APIs
1. Dockerize protoc tools (protoc-gen-grpc-gateway, protoc-gen-openapiv2, protoc-gen-go, protoc-gen-go-grpc)
