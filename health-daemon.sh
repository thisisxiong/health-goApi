#! /bin/bash
port=${1:-8080}
while true; do
    num=$(ps -axu | grep -w health | grep -v grep | wc -l)
    if [ "$num" -lt 1 ]; then
        ./health -port="$port" > log/app_$(date +"%Y%m%d").log 2>&1
    fi
    sleep 10
done
