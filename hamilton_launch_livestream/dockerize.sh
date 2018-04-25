DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
IMAGE_NAME="soar/hamilton_launch_livestream"
DOCKER_HOME="/hamilton_launch_livestream/"

if [ "$1" == "init" ]; then
    docker build -t $IMAGE_NAME .
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        $IMAGE_NAME ffserver
elif [ "$1" == "test" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" \
        $IMAGE_NAME ls -al
elif [ "$1" == "run" ]; then
    docker run --rm \
        --volume "$DIR:$DOCKER_HOME" 
else
    echo "usage: $0 [ init | test | run [args] ]"
fi
