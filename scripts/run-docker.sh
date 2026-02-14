#!/bin/bash

set -euo pipefail

# Cleanup docker containers and exit with correct status
function cleanup() {
    echo "--- stopping existing containers"
    docker stop "knitary"
    docker rm "knitary"

    echo "--- removing knitary images"
    docker rmi "knitary-image"
}
# run cleanup upon terminal exit
trap cleanup EXIT

docker build -t "knitary-image" ../

docker run --name "knitary" knitary-image
