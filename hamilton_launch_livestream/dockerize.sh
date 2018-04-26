DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_livestream"
DOCKER_HOME="/hamilton_launch_livestream/"


test_count() {
    local num_devices=`ls /dev/ | grep "video" | wc -l`
    local device_string=""
    local ffmpeg_command=""

    docker run -d --rm \
        --volume "$DIR:$DOCKER_HOME" \
        -p 8090:8090 \
        $IMAGE_NAME ffserver -f ffserver.conf

    for (( i=0; i<=$((num_devices-1)); i++)); do
        device_string=$device_string"--device=/dev/video$((i)) "
        #ffmpeg_command=$ffmpeg_command"ffmpeg -f video4linux2 -s 640x480 -r 30 \
            #-input_format mjpeg -i /dev/video$((i)) http://localhost:8090/feed$((i+1)).ffm \
             #-nostdin -nostats"
        #${ffmpeg_command}

        docker run --rm \
            --volume "$DIR:$DOCKER_HOME" \
            -p 8090:809$((i)) \
            --device=/dev/video$((i)) \
            $IMAGE_NAME ffmpeg -f video4linux2 -s 640x480 -r 30 \
                -input_format mjpeg \
                -i /dev/video$((i)) \
                http://localhost:8090/feed$((i+1)).ffm
    done
}


if [ "$1" == "init" ]; then
    docker build -t $IMAGE_NAME .
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        $IMAGE_NAME ffserver --help
elif [ "$1" == "test-count" ]; then
    test_count
elif [ "$1" == "ffserver" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        $IMAGE_NAME ffserver --help
elif [ "$1" == "run" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" 
else
    echo "usage: $0 [ init | test | run [args] ]"
fi
