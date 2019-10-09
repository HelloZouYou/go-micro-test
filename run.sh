#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "DIR：${DIR}"
# 构建 auth
cd  $DIR/auth
GOOS=linux go build -o auth *.go
# 构建 inventory-srv
cd  $DIR/inventory-srv
GOOS=linux go build -o inventory-srv *.go
# 构建 orders-srv
cd  $DIR/orders-srv
GOOS=linux go build -o orders-srv *.go
# 构建 orders-web
cd  $DIR/orders-web
GOOS=linux go build -o orders-web *.go
# 构建 payment-srv
cd  $DIR/payment-srv
GOOS=linux go build -o payment-srv *.go
# 构建 payment-web
cd  $DIR/payment-web
GOOS=linux go build -o payment-web *.go
# 构建 user-srv
cd  $DIR/user-srv
GOOS=linux go build -o user-srv *.go
# 构建 user-web
cd  $DIR/user-web
GOOS=linux go build -o user-web *.go

cd $DIR
docker-compose up -d --build auth inventory-srv orders-srv orders-web payment-srv payment-web user-srv user-web