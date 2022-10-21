#!/bin/bash

flag=$1
function daemonStart() {
  while true
  do
    num=`ps axu|grep health|grep -v grep|wc -l`
    if [ $num -lt 1 ];then
      nohup ./health >> start.log 2>&1 &
    fi
    sleep 10
  done
}

function build() {
  `go build -o health`
}


if [ $flag == "start" ]
then
  daemonStart
elif [ $flag == "build" ]
then
  build
fi





