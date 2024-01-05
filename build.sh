#!/bin/bash

# Builds the docker remote controll docker image.

OPTIND=1         # Reset in case getopts has been used previously in the shell.

# Initialize variables

# Image name and Tag
TAG=latest
USER_NAME=$(whoami)
USER_ID=$(id -u)
GROUP_ID=$(id -g)


help() {
	echo "Usage: $(basename $0) [\
		-t <Tag (default: $TAG)> \
		-u <user id (default: $UID)> \
		-g <group id (default: $GID)> \
		-n <user name (default: $UNAME)>] " 	
}

while getopts "h?g:u:t:n:" opt; do
  case "$opt" in
    h|\?)
	usage; exit 0
      ;;
    t)  TAG=$OPTARG
      ;;
	u)
		USER_ID=$OPTARG
		;;
	g)
		GROUP_ID=$OPTARG
		;;
	n)
		USER_NAME=$OPTARG
		;;


  esac
done
shift $((OPTIND-1))
[ "${1:-}" = "--" ] && shift

#echo "$USER_NAME ($USER_ID:$GROUP_ID) -> $TAG"

docker build --build-arg UID=$USER_ID --build-arg GID=$GROUP_ID --build-arg USER_NAME=$UNAME -t remotecontrol:$TAG .
