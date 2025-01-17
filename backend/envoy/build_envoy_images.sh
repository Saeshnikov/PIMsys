#!/bin/sh

docker build -t envoy/sso ./sso
docker build -t envoy/shop ./shop
docker build -t envoy/branch ./branch
docker build -t envoy/template ./template
