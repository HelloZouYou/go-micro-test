#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# 构建 user-srv
cd  $DIR/user-srv
GOOS=linux go build -o user-srv *.go

cd $DIR
docker-compose down
docker-compose up -d