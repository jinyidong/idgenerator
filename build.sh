#!/bin/bash
echo ---------------build main.go...------------------
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -a -installsuffix cgo -o ./IdGenerator ./main.go