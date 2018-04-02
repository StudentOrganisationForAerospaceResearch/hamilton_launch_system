#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_server"
DOCKER_HOME="/hamilton_launch_server/"

ACTION="$1"

UNAME_OUT="$(uname -s)"
case "${UNAME_OUT}" in
    Linux*)     PLATFORM="lin";;
    Darwin*)    PLATFORM="osx";;
    CYGWIN*)    PLATFORM="win";;
    MINGW*)     PLATFORM="win";;
    *)          PLATFORM="unknown"
esac

if [ "$PLATFORM" == "unknown" ]; then
    echo "WARNING: Could not detect platform"
else
    echo Detected "$PLATFORM" system
fi

build_linux() {
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        --env GOOS=linux \
        --env GOARCH=amd64 \
        $IMAGE_NAME go build
}

build_windows() {
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        --env GOOS=windows \
        --env GOARCH=amd64 \
        $IMAGE_NAME go build
}

build_osx() {
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        --env GOOS=darwin \
        --env GOARCH=amd64 \
        $IMAGE_NAME go build
}

build_rasperry_pi() {
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        --env GOOS=linux \
        --env GOARCH=arm \
        --env GOARM=5 \
        $IMAGE_NAME go build
}


if [ "$ACTION" == "init" ]; then
    docker build -t $IMAGE_NAME .
elif [ "$ACTION" == "build" ]; then
    target="$2"
    if [ -z "$target" ]; then
        target="$PLATFORM"
    fi

    if [ "$target" == "lin" ]; then
        echo "Building..."
        build_linux
        echo "Done"
    elif [ "$target" == "win" ]; then
        echo "Building..."
        build_windows
        echo "Done"
    elif [ "$target" == "osx" ]; then
        echo "Building..."
        build_osx
        echo "Done"
    elif [ "$target" == "rpi" ]; then
        echo "Building..."
        build_rasperry_pi
        echo "Done"
    else
        echo "ERROR: Unknown target platform [ lin | osx | win | rpi ]"
    fi
else
    echo "usage: $0 init "
    echo "       $0 build [ lin | osx | win | rpi ]"
    exit 1
fi
