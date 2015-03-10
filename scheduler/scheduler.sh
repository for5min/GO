#!/bin/bash
PROG=/home/scaleworks/scheduler
CONF=/home/scaleworks/scheduler.json
LOGFILE=/home/scaleworks/${STAGING}/current/log/scheduler.log
PIDFILE=/var/tmp/scheduler.pid
PID=`cat /var/tmp/scheduler.pid`

case $1 in
  start)
    $PROG $CONF >> $LOGFILE 2>&1
    echo $! > $PIDFILE
    ;;
  stop)
    kill -9 $PID
    ;;
  *) echo "start|stop"
    ;;
esac
