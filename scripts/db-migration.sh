#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../

cd ./schema/db

# $1=up or down
# $2=local or localtest
sql-migrate "$1" -env="$2"
