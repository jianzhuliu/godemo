## cluster 
>- mkdir -p /app/redis6/conf/cluster

```
## 设置加入 cluster
cluster-enabled yes 

## cluster 配置文件名
cluster-config-file nodes-6379

## 节点超时，判断下线或者切换节点 master replica, 单位毫秒
cluster-node-timeout 10000
```

## 准备配置文件
>- /app/redis6/conf/cluster/redis-6379.conf 
```
include /app/redis6/conf/comm.conf

port 6379
logfile "6379.log"
dbfilename dump-6379.rdb
appendfilename "appendonly-6379.aof"

cluster-enabled yes 
cluster-config-file nodes-6379.conf 

##10秒超时时间
cluster-node-timeout 10000 

```

>- cd /app/redis6/conf/cluster
>- sed "s/6379/6380/g" redis-6379.conf > redis-6380.conf
>- sed "s/6379/6381/g" redis-6379.conf > redis-6381.conf
>- sed "s/6379/6382/g" redis-6379.conf > redis-6382.conf
>- sed "s/6379/6383/g" redis-6379.conf > redis-6383.conf
>- sed "s/6379/6384/g" redis-6379.conf > redis-6384.conf
>- sed "s/6379/6384/g" redis-6379.conf > redis-6385.conf

## 启动 master replica 
```
redis-server /app/redis6/conf/cluster/redis-6380.conf
redis-server /app/redis6/conf/cluster/redis-6381.conf
redis-server /app/redis6/conf/cluster/redis-6382.conf
redis-server /app/redis6/conf/cluster/redis-6383.conf
redis-server /app/redis6/conf/cluster/redis-6384.conf
redis-server /app/redis6/conf/cluster/redis-6385.conf

ps -ef | grep redis 
```


## 配置集群
```
redis-cli --cluster create --cluster-replicas 1 \
127.0.0.1:6380 \
127.0.0.1:6381 \
127.0.0.1:6382 \
127.0.0.1:6383 \
127.0.0.1:6384 \
127.0.0.1:6385

```

## 操作 
>- redis-cli -p 6380
```
set name jianzhu
```

>- redis-cli -c -p 6380
```
set name jianzhu
```

## 关闭 replica  
>- 查看集群对应关系  m 6381 -> s 6383
```
redis-cli -c -p 6380 
cluster nodes
```

>- 关闭 6383
```
redis-cli -c -p 6383
shutdown 

```


>- 查看日志  
```
tail -f /app/redis6/data/6381.log 
tail -f /app/redis6/data/6383.log 

```

>- 启动 6383
```
redis-server /app/redis6/conf/cluster/redis-6383.conf
```

## 关闭 master 
>- 关闭 6381
```
ps -ef | grep redis | grep 6381 | awk '{print $2}' | xargs kill -9 

```

>- 查看日志 
```
tail -f /app/redis6/data/6381.log 
tail -f /app/redis6/data/6383.log 

```

>- 启动 6381
```
redis-server /app/redis6/conf/cluster/redis-6381.conf
```

## 清理
>- ps -ef | grep redis | grep -v "grep" | awk '{print $2}' | xargs kill -9 






