DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_livestream"
DOCKER_HOME="/hamilton_launch_livestream/"


test_count() {
    local num_devices=`ls /dev/ | grep "video" | wc -l`

    # devices start from /dev/video0
    # target url starts with /feed1
    for (( i=0; i<=$((num_devices-1)); i++)); do
        device_string="--device=/dev/video$((i)) "
        ffmpeg_command="ffmpeg -f video4linux2 -s 640x480 -r 30 \
            -input_format mjpeg -i /dev/video$((i)) http://ffserver.livestream-net:8090/feed$((i+1)).ffm\
            -nostdin -nostats"

        docker run --rm \
            --volume "$DIR:$DOCKER_HOME" \
            --network livestream-net \
            ${device_string} \
            $IMAGE_NAME ${ffmpeg_command}
    done
}

launch_ffserver() {
    docker run --rm \
        --name ffserver \
        --volume "$DIR:$DOCKER_HOME" \
        --network livestream-net \
        -p 8090:8090 \
        $IMAGE_NAME ffserver -f ffserver.conf
}

if [ "$1" == "init" ]; then
    docker build -t $IMAGE_NAME .
elif [ "$1" == "ffmpeg" ]; then
    test_count
elif [ "$1" == "ffserver" ]; then
    launch_ffserver
else
    echo "usage: $0 [ init | ffmpeg | ffserver ]"
fi
