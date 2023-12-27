#!/bin/bash

ldflags="-s -w"

go env -w GOPROXY=https://goproxy.cn,direct

# Plugin
CGO_ENABLED=0 GOARCH=$(go env GOARCH) GOOS=$(go env GOOS) go build -ldflags "$ldflags" -o plugin/bin/plugin plugin/plugin.go
CGO_ENABLED=0 GOARCH=$(go env GOARCH) GOOS=$(go env GOOS) go build -ldflags "$ldflags" -o plugin/bin/example plugin/test/example.go

# Replica
CGO_ENABLED=0 GOARCH=$(go env GOARCH) GOOS=$(go env GOOS) go build -ldflags "$ldflags" -o replica/bin/replica replica/replica.go
CGO_ENABLED=0 GOARCH=$(go env GOARCH) GOOS=$(go env GOOS) go build -ldflags "$ldflags" -o replica/bin/example replica/test/example.go
