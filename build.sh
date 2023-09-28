#!/usr/bin/env bash
# Stops the process if something fails
# set -xe
# touch /var/app/current/go.bak
# sudo rm /var/app/current/go.*

# go get -u

# GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"

# chmod +x bin/application

go get -u

go build -o bin/application