#!/bin/bash
GOOS=darwin GOARCH=amd64 go build -o mac.ipynb *go
GOOS=windows GOARCH=amd64 go build -o windows.ipynb
GOOS=linux GOARCH=amd64 go build -o linux.ipynb
