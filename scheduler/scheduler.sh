#!/bin/bash
<<<<<<< HEAD
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
=======
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
>>>>>>> 697fa299018f9db9e5341aab4bceb78793642bf2
