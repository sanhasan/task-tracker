#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

APP_NAME="todosher"
CMD_PATH="./cmd/app"
OUT_DIR="pkg"

mkdir -p "${OUT_DIR}"

echo "Building ${APP_NAME} for multiple platforms..."

build() {
    local os=$1
    local arch=$2
    local ext=""

    if [ "$os" = "windows" ]; then
        ext=".exe"
    fi

    local out_name="${APP_NAME}-${os}-${arch}${ext}"
    echo "   Building ${os}/${arch} -> ${OUT_DIR}/${out_name}"

    GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -o "${OUT_DIR}/${out_name}" "${CMD_PATH}"
}

build linux amd64
build linux arm64
build darwin amd64
build darwin arm64
build windows amd64

echo "OK: All builds completed successfully in '${OUT_DIR}/' directory!"