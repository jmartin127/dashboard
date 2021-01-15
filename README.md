# dashboard

This is an example project I put together to demonstrate the utility of some awesome services that exist within the gRPC ecosystem.

# References

gRPC Resources: https://github.com/grpc-ecosystem/awesome-grpc#documentation

Official Website: https://grpc.io/

Helpful Quickstart: https://grpc.io/docs/languages/go/quickstart/

gRPCurl: https://github.com/fullstorydev/grpcurl

grpc-json-proxy: https://github.com/jnewmano/grpc-json-proxy

# Generating code from Proto

```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    traffic/traffic.proto
```

```
$ protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     weather/weather.proto
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
