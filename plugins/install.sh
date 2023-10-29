#!/bin/bash
go build -buildmode=plugin -o $(go env GOPATH)/bin/hugo.so
