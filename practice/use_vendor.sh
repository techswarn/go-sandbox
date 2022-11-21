#!/usr/bin/env bash

# Cleanup
rm -f hello/app 2> /dev/null

docker run --rm \
  -v "$(pwd)/hello:/mounted" \
  --platform linux/amd64 golang:1.19 \
  /bin/bash -c \
  '
  cd /mounted || exit
  env GOOS=darwin GOARCH=amd64 go build -mod=vendor -o app
  exit 0
  '