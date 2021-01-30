#!/bin/bash

go build -o $PWD/bin/traffic $PWD/cmd/traffic/main.go
go build -o $PWD/bin/weather $PWD/cmd/weather/main.go
go build -o $PWD/bin/dashboard $PWD/cmd/dashboard/main.go
go build -o $PWD/bin/gateway $PWD/cmd/gateway/main.go