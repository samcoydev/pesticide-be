#!/usr/bin/env bash
set -xe

# Install dependencies.
go get ./...

# Build app
go build -o bin/application application.go