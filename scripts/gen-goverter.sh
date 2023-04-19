#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../

cd ./src/backend

go run github.com/jmattheis/goverter/cmd/goverter@latest ./adapter/gateway/convert
