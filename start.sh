#!/bin/bash

flag=$1
addr=${2:-8080}
function start() {
    nohup ./app-daemon.sh "$addr" &
}

function build() {
    $(go build -o health)
}

function status(){
    d=$(ps axu | grep -w app-daemon | grep -v grep | awk '{print $2}')
    dPid=${d:-0}
    if [ "$dPid" -gt 0 ]; then
        echo "进程守护脚本已开启：$dPid"
    else
        echo -e "33[31m 进程守护脚本未开启 33[0m"
    fi

    h=$(ps axu | grep -w health | grep -v grep | awk '{print $2}')
    hPid=${h:-0}
    if [ "$hPid" -gt 0 ]; then
        echo "程序已运行：$hPid"
    else
        echo -e "33[31m 程序未运行 33[0m"
    fi
}

function stop() {
    d=$(ps axu | grep -w app-daemon | grep -v grep | awk '{print $2}')
	dPid=${d:-0}
    if [ "$dPid" -gt 0 ]; then
        kill -9 "$d"
        echo "守护脚本关闭"
    fi
    h=$(ps axu | grep -w health | grep -v grep | awk '{print $2}')
	hPid=${h:-0}
    if [ "$hPid" -gt 0 ]; then
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
*)
    echo -e "请选择你需要执行的动作 :\n start [port] 守护进行方式启动程序 \n stop 关闭守护进程和程序 \n build 重新构建程序"
    ;;
esac
