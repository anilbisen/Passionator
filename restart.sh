#!/bin/bash
if [ x = x$1 ]; then
    echo "Usage: ./restart.sh <service>"
    echo "    <service>  : The service, e.g. passionator"
  exit
fi
SERVICE=$1

docker-compose stop $SERVICE
docker-compose rm -f $SERVICE
docker-compose create $SERVICE
docker-compose start $SERVICE
docker-compose logs $SERVICE
