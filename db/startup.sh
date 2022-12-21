#!/usr/bin/env bash

MYSQL_IMAGE_NAME="mysql8"

# 如果没有镜像则创建镜像
if [ -z "$(docker images -q $MYSQL_IMAGE_NAME)" ]; then
  docker build -t $MYSQL_IMAGE_NAME .
fi

# mysql 没有启动则启动
if [ -z "$(docker ps -q -f name=mysql)" ]; then
  docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root $MYSQL_IMAGE_NAME \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_unicode_ci
  echo "mysql container started"
fi

# phpmyadmin 没有启动则启动
if [ -z "$(docker ps -q -f name=phpmyadmin)" ]; then
  docker run -d --name phpmyadmin -p 8080:80 --link mysql:db -e PMA_HOST=db phpmyadmin
  echo "phpmyadmin container started"
fi
