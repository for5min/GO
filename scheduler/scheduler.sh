#!/bin/bash
PROG=/home/scaleworks/scheduler
CONF=/home/scaleworks/scheduler.json

PIDFILE=/var/tmp/scheduler.pid
PID=`cat /var/tmp/scheduler`

case $1 in
  start)
    $PROG $CONF 2>&1
    echo $! > $PIDFILE
    ;;
  stop)
    kill -9 $PID
    ;;
  restart)
    stop
    start
    ;;
  *) echo "start|stop|restart"
    ;;
  esac
