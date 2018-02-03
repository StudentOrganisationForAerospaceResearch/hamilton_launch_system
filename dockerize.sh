#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_system"
CONTAINER_NAME="hamilton_launch_system"
DOCKER_HOME="/hamilton_launch_system/"

if [ "$1" == "init" ]; then
    docker build -t $IMAGE_NAME $DIR
elif [ "$1" == "run" ]; then
    docker kill "$CONTAINER_NAME"
    docker run --rm \
        --name "$CONTAINER_NAME" \
        --volume "$DIR":"$DOCKER_HOME" \
        -p "${2}":"${2}" \
        "$IMAGE_NAME" python3 -u src/main.py "${2}"
elif [ "$1" == "shell" ]; then
    docker exec -it "$CONTAINER_NAME" /bin/bash
else
    echo "usage: $0 [ init | run [port] | shell ]"
    exit 1
fi
