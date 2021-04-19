#!/bin/bash

####
## 单例运行脚本
####

####最简单的方式,一行代码搞定
#[ "${FLOCKER}" != "$0" ] && exec env FLOCKER="$0" flock -en "$0" "$0" "$@" || :

#sleep 10 
#echo "done"
#exit 1

##锁文件
LOCK_FILE=/tmp/flock.lock

##创建文件描述符 99，并指向这个锁文件
exec 99>"${LOCK_FILE}"

## 尝试非阻塞模式加锁
flock -n 99

## 如果加锁失败，表示已经有实例在运行
if [ "$?" != 0 ];then
   echo "$0 is already running"
   exit 1
fi

sleep 10
echo "done"


