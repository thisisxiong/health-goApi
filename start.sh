#!/bin/bash

flag=$1
addr=${2:-8080}
function start() {
    nohup ./health-daemonStart.sh "$addr" &
}

function build() {
    $(go build -o health)
}

function stop() {
    d=$(ps axu | grep -w health-daemon | grep -v grep | awk '{print $2}')
    if [ "$d" -gt 0 ]; then
        kill -9 "$d"
        echo "守护脚本关闭"
    fi
    h=$(ps axu | grep -w health | grep -v grep | awk '{print $2}')
    if [ "$h" -gt 0 ]; then
        kill -9 "$h"
        echo "程序关闭"
    fi
}

case $flag in
"start")
    start
    ;;
"build")
    build
    ;;
"stop")
    stop
    ;;
esac
