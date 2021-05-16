## redis
>- https://redis.io/
>- http://redis.cn/
>- https://hub.docker.com/_/redis/
>- https://github.com/gomodule/redigo

## install
```
wget https://download.redis.io/releases/redis-6.2.2.tar.gz
tar -zxf redis-6.2.2.tar.gz
cd redis-6.2.2
yum install -y gcc 
make USE_SYSTEMD=yes
make PREFIX=/app/redis6 install

cat >>/etc/profile<<EOF
REDIS_HOME=/app/redis6
PATH=\$PATH:\$REDIS_HOME/bin
EOF

source /etc/profile

```

## docker start

```
docker pull redis:6.2.2
docker run --name=redis6 -d -p 6379:6379 redis:6.2.2
docker exec -it redis6 redis-cli 
```


## 常用类型 
>- string
>- hash
>- list
>- set
>- sorted_set

## key 操作
>- set k1 v1 
>- hset hk1 f1 v1 
>- lpush lk1 v1 v2 v3 
>- sadd sk1 v1 v2 
>- zsadd zk1 1 v1 2 v2 
 
>- type k1
>- exists k1
>- del k1


## string
>- [string](./README.string.md)

## hash
>- [hash](./README.hash.md)

## list 
>- [list](./README.list.md)

## set 
>- [set](./README.set.md)

## sorted_set
>- [sorted_set](./README.sorted_set.md)

## 哨兵 sentinel
>- [sentinel](./README.sentinel.md)

## 集群 cluster 
>- [cluster](./README.cluster.md)
