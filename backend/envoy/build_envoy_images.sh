#!/bin/sh

docker build -t envoy/sso ./sso
docker build -t envoy/shop ./shop