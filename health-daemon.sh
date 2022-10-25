#! /bin/bash
port=${1:-8080}
while true; do
    num=$(ps -axu | grep -w health | grep -v grep | wc -l)
    if [ "$num" -lt 1 ]; then
        ./health -port="$port" > start.log 2>&1
    fi
    sleep 10
done
