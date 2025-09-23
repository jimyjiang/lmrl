#!/bin/bash

# 定义参数
WORK_DIR="/root/deploy/lmrl"
NAME="lmrl"
CMD="./lmrl"
PID_FILE="/var/run/${NAME}.pid"
# LOG_FILE="/var/log/${NAME}.log"

cd "${WORK_DIR}" || exit 1

# 启动函数
start() {
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        if ps -p $PID > /dev/null; then
            echo "$NAME is already running (PID $PID)"
            return 1
        else
            echo "Removing stale PID file"
            rm -f "$PID_FILE"
        fi
    fi
    
    echo "Starting $NAME..."
    $CMD >> /dev/null 2>&1 &
    PID=$!
    echo $PID > "$PID_FILE"
    echo "Started $NAME with PID $PID"
}

# 停止函数
stop() {
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        if ps -p $PID > /dev/null; then
            echo "Stopping $NAME (PID $PID)..."
            kill $PID
            rm -f "$PID_FILE"
            echo "$NAME stopped"
        else
            echo "$NAME is not running (stale PID file)"
            rm -f "$PID_FILE"
        fi
    else
        echo "$NAME is not running (no PID file)"
    fi
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    *)
        echo "Usage: $0 {start|stop|restart}"
        exit 1
        ;;
esac