#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../

cd ./src/backend
rm -f ./adapter/controller/generated.go
rm -f ./adapter/controller/models_gen.go

# https://gqlgen.com/
# https://github.com/99designs/gqlgen
#go get -u github.com/99designs/gqlgen@v0.13.0
gqlgen generate
