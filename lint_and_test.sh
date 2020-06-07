#!/usr/bin/env bash

export GOOGLE_APPLICATION_CREDENTIALS="/home/me/workspace_learn/serverless_go_firebase/creds/service_account.json"

golangci-lint run --no-config --issues-exit-code=1 --enable-all --disable=godot --disable=gocyclo --disable=gochecknoinits --disable=nakedret --disable=gochecknoglobals --tests=false --disable=goimports --disable=wsl \
 --skip-dirs "(^|/)templates($|/)"


# godotenv loads teh env.yml file before running tests
# install it like this: go get github.com/joho/godotenv/cmd/godotenv
# count=1 disables the test cache
godotenv -f .env.yml go test -v ./... -count=1 --cover