#!/bin/bash

trap 'kill %1; kill %2; kill %3; kill %4' SIGINT
$PWD/bin/traffic & $PWD/bin/weather & $PWD/bin/dashboard & $PWD/bin/user  & SWAGGER_UI_DIR=/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/dashboard/swaggerui $PWD/bin/gateway
