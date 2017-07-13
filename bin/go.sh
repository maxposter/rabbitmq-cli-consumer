#!/bin/sh
set -e
cd $(cd $(dirname $0)/../ && pwd)

cmd=$@

docker-compose run --rm go bash -c "${cmd}"