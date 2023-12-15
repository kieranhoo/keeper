#!/usr/bin/env bash

go test -v ./pkg/backend
go test -v ./pkg/brokers
go test -v ./pkg/tasks
go test -v ./pkg/worker

