all: gen swagger build run

gen:
	cd cmd/traffic/ && make gen
	cd cmd/weather/ && make gen

swagger:
	cd proto && \
	protoc -I . \
     --openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=dashboard.json \
     jmartin127/traffic/v1/traffic.proto jmartin127/weather/v1/weather.proto && \
	cp gen/openapiv2/dashboard.swagger.json ../swaggerui/

build:
	./scripts/build.sh

run:
	./scripts/run.sh
