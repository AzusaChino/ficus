#!bin/sh

# load .env variables
eval "$(echo $(cat .env))"

IMAGE_NAME="$AUTHOR/$PROJECT:$VERSION"

docker build -t $IMAGE_NAME
docker push $IMAGE_NAME
