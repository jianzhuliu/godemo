## 哨兵 sentinel

## 1master + 2replica + 3sentinel
>- 配置文件存放目录 /app/redis6/conf 
>- 数据存放目录 /app/redis6/data

#### master
>- ./conf/redis-6379.conf

```
include /app/redis6/conf/comm.conf

port 6379
#logfile "6379.log"
dbfilename dump-6379.rdb
appendfilename "appendonly-6379.aof"
```

#### replica
> ./conf/redis-6380.conf 
```
include /app/redis6/conf/comm.conf

port 6380
#logfile "6380.log"
dbfilename "dump-6380.rdb"
appendfilename "appendonly-6380.aof"

replicaof 127.0.0.1 6379
```

> ./conf/redis-6381.conf 
```
include /app/redis6/conf/comm.conf

port 6381
#logfile "6381.log"
dbfilename "dump-6381.rdb"
appendfilename "appendonly-6381.aof"

replicaof 127.0.0.1 6379
```

#### sentinel
>- ./conf/sentinel-26379.conf
```
port 26379
daemonize no
#pidfile /var/run/redis-sentinel.pid
#logfile ""
dir "/app/redis6/data"
sentinel monitor mymaster 127.0.0.1 6379 2
sentinel down-after-milliseconds mymaster 30000
sentinel parallel-syncs mymaster 1
sentinel failover-timeout mymaster 180000
sentinel deny-scripts-reconfig yes

```

>- ./conf/sentinel-26380.conf
```
port 26380
daemonize no
#pidfile /var/run/redis-sentinel.pid
#logfile ""
dir "/app/redis6/data"
sentinel monitor mymaster 127.0.0.1 6379 2
sentinel down-after-milliseconds mymaster 30000
sentinel parallel-syncs mymaster 1
sentinel failover-timeout mymaster 180000
sentinel deny-scripts-reconfig yes

```

>- ./conf/sentinel-26381.conf
```
port 26381
daemonize no
#pidfile /var/run/redis-sentinel.pid
#logfile ""
dir "/app/redis6/data"
sentinel monitor mymaster 127.0.0.1 6379 2
sentinel down-after-milliseconds mymaster 30000
sentinel parallel-syncs mymaster 1
sentinel failover-timeout mymaster 180000
sentinel deny-scripts-reconfig yes

```

## 启动顺序
>- 启动 master 
```
redis-server /app/redis6/conf/redis-6379.conf
```

>- 逐步启动2个 replica
```
redis-server /app/redis6/conf/redis-6380.conf
redis-server /app/redis6/conf/redis-6381.conf
```

>- 逐步启动3个 sentinel
```
redis-sentinel /app/redis6/conf/sentinel-26379.conf
redis-sentinel /app/redis6/conf/sentinel-26380.conf
redis-sentinel /app/redis6/conf/sentinel-26381.conf

```



