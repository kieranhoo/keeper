#!/usr/bin/env bash

go test -v .

go test -v ./internal/logs
go test -v ./internal/errors
go test -v ./internal/core
