#!/usr/bin/bash

# load .env variables
eval "$(echo $(cat .env))"

PUSH=$1

IMAGE_NAME="$AUTHOR/$PROJECT:$VERSION"
echo "building image $IMAGE_NAME, please wait..."

docker build . -t $IMAGE_NAME
if [ $PUSH ]; then
  docker push $IMAGE_NAME
fi
