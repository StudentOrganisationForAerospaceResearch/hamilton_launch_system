#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_board"
DOCKER_HOME="/hamilton_launch_board/"

if [ "$1" == "init" ]; then
    docker build -t $IMAGE_NAME .
elif [ "$1" == "install" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        $IMAGE_NAME yarn install --frozen-lockfile
elif [ "$1" == "start" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        -p 8080:8080 \
        $IMAGE_NAME npm start
elif [ "$1" == "run" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        $IMAGE_NAME npm run "${@:2}"
else
    echo "usage: $0 [ init | install | start | run [args] ]"
fi
