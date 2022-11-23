#!/bin/zsh
go mod tidy
go build -o app
GODAILYLIB_ENV=production ./app