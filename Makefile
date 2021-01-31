all: gen swagger build run

gen:
	cd cmd/traffic/ && make gen
	cd cmd/weather/ && make gen
	cd cmd/dashboard/ && make gen
	cd cmd/user/ && make gen

swagger:
	cd proto && \
	protoc -I . \
     --openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=dashboard.json \
     jmartin127/traffic/v1/traffic.proto jmartin127/weather/v1/weather.proto jmartin127/dashboard/v1/dashboard.proto jmartin127/user/v1/user.proto && \
	cp gen/openapiv2/dashboard.swagger.json ../swaggerui/

build:
	./scripts/build.sh

run:
	./scripts/run.sh
