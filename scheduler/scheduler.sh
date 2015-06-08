#!/bin/bash
RELEASE=/home/scaleworks/production/current

PROG=/usr/bin/scheduler
CONF=/usr/bin/scheduler.json


LOGFILE=${RELEASE}/log/scheduler.log
PIDFILE=${RELEASE}/tmp/pids/scheduler.pid



case $1 in
  start)
    $PROG $CONF >> $LOGFILE 2>&1 &
    echo $! > $PIDFILE
    ;;
  stop)
    PID=`cat ${PIDFILE}`
    kill 9 $PID
     rm $PIDFILE
    ;;
  restart)
    stop
    start
    ;;
  *) echo "start|stop|restart"
    ;;
  esac
