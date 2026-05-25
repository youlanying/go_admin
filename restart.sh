#!/bin/bash
set -e
echo "go build"
go mod tidy
CGO_ENABLED=1 go build -tags sqlite3 -o go-admin main.go
chmod +x ./go-admin
echo "kill go-admin service"
killall go-admin 2>/dev/null || true
echo "run go-admin"
exec ./go-admin server -c=config/settings.sqlite.yml
