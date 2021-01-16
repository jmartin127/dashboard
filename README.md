# dashboard

This is an example project I put together to demonstrate the utility of some awesome services that exist within the gRPC ecosystem.

# References

gRPC Resources: https://github.com/grpc-ecosystem/awesome-grpc#documentation

Official Website: https://grpc.io/

Helpful Quickstart: https://grpc.io/docs/languages/go/quickstart/

gRPCurl: https://github.com/fullstorydev/grpcurl

grpc-json-proxy: https://github.com/jnewmano/grpc-json-proxy

grpc-gateway: https://github.com/grpc-ecosystem/grpc-gateway

grpc google.api.http annotations: https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46

all the crazy annotations: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/internal/proto/examplepb/a_bit_of_everything.proto

Buff Style Guide: https://docs.buf.build/style-guide/#files-and-packages

# Generating code from Proto

```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    traffic/traffic.proto
```

Generate ALL:
```
   protoc -I . \
     --go_out=./gen/go/ --go_opt=paths=source_relative \
     --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative \
     --grpc-gateway_out=./gen/go/ --grpc-gateway_opt paths=source_relative \
     --grpc-gateway_opt logtostderr=true \
     traffic/traffic.proto
```

# Making requests

List out gRPC functions supported by a service:
```
$ grpcurl -plaintext localhost:50051 list traffic.Traffic
traffic.Traffic.GetTravelTime
```

Retrieve Travel Time:
```
$ grpcurl -d '{"origin": "my house", "destination": "Lehi"}' -plaintext localhost:50051 traffic.Traffic/GetTravelTime
{
  "travelTime": "660s"
}
```

Retrieve Current Weather:
```
$ grpcurl -d '{"address": "lehi"}' -plaintext localhost:50052 weather.Weather/GetCurrentWeather
{
  "tempFahrenheit": 39,
  "precipitationPct": 1,
  "humidityPct": 38,
  "windMPH": 5
}
```

Retrieve travel time via REST:
```
$ curl -X POST "localhost:8081/traffic/travel/time"
{"travelTime":"660s"}
```

Retrieve weather via REST:
```
$ curl -X POST "localhost:8081/weather/current"
{"tempFahrenheit":39,"precipitationPct":1,"humidityPct":38,"windMPH":5}
```

# TODO
1. Actually doing GET instead of POSTs in google.api annotations
1. Working Swagger (Open API) pages
1. Lyft validation
1. Dockerize protoc tools (protoc-gen-grpc-gateway, protoc-gen-openapiv2, protoc-gen-go, protoc-gen-go-grpc)
1. Launch scripts
