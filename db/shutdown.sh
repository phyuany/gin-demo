#!/usr/bin/env bash

MYSQL_IMAGE_NAME="mysql8"

# shutdown mysql container if it is running
if [ -n "$(docker ps -a -q -f name=mysql)" ]; then
  docker stop mysql
  docker rm mysql
  echo "mysql container stopped and removed"
fi

# shutdown phpmyadmin container if it is running
if [ -n "$(docker ps -a -q -f name=phpmyadmin)" ]; then
  docker stop phpmyadmin
  docker rm phpmyadmin
  echo "phpmyadmin container stopped and removed"
fi

# remove mysql image if it exists
if [ -n "$(docker images -q $MYSQL_IMAGE_NAME)" ]; then
  docker rmi $MYSQL_IMAGE_NAME
  echo "mysql image removed"
fi
