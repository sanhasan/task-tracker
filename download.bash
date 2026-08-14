#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

APP_NAME="todosher"
CMD_PATH="./cmd/app/${APP_NAME}"
OUT_DIR="pkg"

mkdir -p "${OUT_DIR}"

echo "Building ${APP_NAME} from ${CMD_PATH}..."
go build -o "${OUT_DIR}/${APP_NAME}" "${CMD_PATH}"

echo "OK: ${OUT_DIR}/${APP_NAME}"