#!/usr/bin/env bash

set -e
cd $(cd $(dirname $0)/../ && pwd)

echo "Removing all services...";
docker-compose rm -fv

echo "Building all services...";
docker-compose build

echo "Building application ..."
mkdir -p output
./bin/go.sh "cd /go/src/github.com/maxposter/rabbitmq-cli-consumer && go get && go build -o output/rabbitmq-cli-consumer"
