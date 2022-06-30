#!/bin/sh

go mod tidy

go build -v -o ./user/main ./user/rpc/user.go
go build -v -o ./videomgr/main ./videomgr/rpc/videomanager.go
go build -v -o ./social/main ./social/rpc/social.go
go build -v -o ./api/main ./api/main.go
