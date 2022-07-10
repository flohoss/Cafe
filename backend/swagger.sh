#!/bin/sh

action=$1

case $action in
"install")
    go install github.com/swaggo/swag/cmd/swag@latest
    ;;
"init")
    swag init --parseDependency -g api/swagger.go
    ;;
*)
    exit 0
    ;;
esac