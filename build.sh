#!/usr/bin/env bash

# Install dependencies.
go get ./...

# Build app
go build ./ -o bin/application