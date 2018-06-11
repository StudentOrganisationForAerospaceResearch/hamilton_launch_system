#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_server"
DOCKER_HOME="/hamilton_launch_server/"

LIVESTREAM_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd livestream && pwd )"
LIVESTREAM_IMAGE_NAME="soar/hamilton_launch_livestream"
LIVESTREAM_DOCKER_HOME="/hamilton_launch_livestream/"

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

run_ffserver() {
    (cd $LIVESTREAM_DIR && 
    docker run --rm \
        --name ffserver \
        --volume "$LIVESTREAM_DIR:$LIVESTREAM_DOCKER_HOME" \
        --network livestream-net \
        -p 8090:8090 \
        $LIVESTREAM_IMAGE_NAME ffserver -f ffserver.conf
    )
}

run_ffmpeg() {
    cd $LIVESTREAM_DIR
    local num_devices=`ls /dev/ | grep "video" | wc -l`

    # devices start from /dev/video0
    # target url starts with /feed1
    for (( i=0; i<=$((num_devices-1)); i++)); do
        device_string="--device=/dev/video$((i)) "
        ffmpeg_command="ffmpeg -f video4linux2 -s 640x480 -r 30 \
            -input_format mjpeg -i /dev/video$((i)) http://ffserver.livestream-net:8090/feed$((i+1)).ffm\
            -nostdin -nostats"

        docker run --rm \
            --name ffmpeg$((i)) \
            --volume "$LIVESTREAM_DIR:$LIVESTREAM_DOCKER_HOME" \
            --network livestream-net \
            ${device_string} \
            $LIVESTREAM_IMAGE_NAME ${ffmpeg_command} &
    done
    cd ..
}

if [ "$ACTION" == "init" ]; then
    echo "Initializing Server..."
    docker build -t $IMAGE_NAME .
    echo "Done"
    echo "Initializing Livestream..."
    (cd livestream &&
    docker build -t $LIVESTREAM_IMAGE_NAME . &&
    docker network create livestream-net
    )
    echo "Done"
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
elif [ "$ACTION" == "ffserver" ]; then
    if [ "$PLATFORM" == "lin" ]; then
        run_ffserver
    else
        echo "ERROR: ffserver command only available on linux"
    fi
elif [ "$ACTION" == "ffmpeg" ]; then
    if [ "$PLATFORM" == "lin" ]; then
        run_ffmpeg
    else
        echo "ERROR: ffmpeg command only available on linux"
    fi
elif [ "$ACTION" == "killall" ]; then
    # Kill the ffmpeg commands first for a clean exit
    docker kill $(docker ps -q -f name=ffmpeg)
    docker kill $(docker ps -q -f name=ffserver)
else
    echo "usage: $0 init"
    echo "       $0 build [ lin | osx | win | rpi ]"
    echo "       $0 ffserver"
    echo "       $0 ffmpeg"
    echo "       $0 killall - Will kill both the ffmpeg and ffserver docker containers"
    exit 1
fi
