#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_system"
CONTAINER_NAME="hamilton_launch_system"
DOCKER_HOME="/hamilton_launch_system/"

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

lint() {
    docker run --rm \
        --volume "$DIR":"$DOCKER_HOME" \
        "$IMAGE_NAME" pyflakes src/*.py
    LINT_RESULT="$?"
    if [ "$LINT_RESULT" -eq 0 ]; then
        printf "\n${GREEN}|*******| LINTING PASSED |*******|${NC}\n\n"
    else
        printf "\n${RED}|*******| LINTING FAILED |*******|${NC}\n\n"
        exit 1
    fi
}

if [ "$1" == "init" ]; then
    docker build -t $IMAGE_NAME $DIR
elif [ "$1" == "lint" ]; then
    lint
elif [ "$1" == "run" ]; then
    lint
    docker kill "$CONTAINER_NAME"
    docker run --rm \
        --name "$CONTAINER_NAME" \
        --volume "$DIR":"$DOCKER_HOME" \
        -p "${2}":"${2}" \
        "$IMAGE_NAME" python3 -u src/main.py "${2}"
elif [ "$1" == "shell" ]; then
    docker exec -it "$CONTAINER_NAME" /bin/bash
else
    echo "usage: $0 [ init | lint | run [port] | shell ]"
    exit 1
fi
