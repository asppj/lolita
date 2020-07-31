#!/bin/bash
docker ps|grep  "lolita"
count=`docker ps |grep  "lolita" | wc -l`
echo "Process Count:" $count
if [ 0 == $count ]; then
echo "Process Start："
docker-compose up -d
else
echo "Process Restart"
docker-compose restart
fi
echo `docker ps|grep "lolita"`