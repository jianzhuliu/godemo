## 缓存

## proxy_cache
>- 存储一些之前被访问过、而且可能将要被再次访问的资源，使用户可以直接从代理服务器获得，从而减少上游服务器的压力，加快整个访问速度

```
语法：proxy_cache zone | off ; # zone 是共享内存的名称
默认值：proxy_cache off;
上下文：http、server、location

```

## proxy_cache_path 
>- 设置缓存文件的存放路径
```
语法：proxy_cache_path path [level=levels] ...可选参数省略
默认值：proxy_cache_path off
上下文：http

```

>- path 缓存文件的存放路径
>- level 目录层级
>- keys_zone 设置共享内存
>- inactive 在指定时间内没有被访问，缓存会被清理，默认10分钟

## proxy_cache_key 
>- 设置缓存文件的 key 
```
语法：proxy_cache_key
默认值：proxy_cache_key $scheme$proxy_host$request_uri;
上下文：http、server、location
```

## proxy_cache_valid 
>- 配置什么状态码可以被缓存，以及缓存时长
```
语法：proxy_cache_valid [code...] time;
上下文：http、server、location
配置示例：proxy_cache_valid 200 304 2m;; # 说明对于状态为200和304的缓存文件的缓存时间是2分钟

```

## proxy_no_cache
>- 定义相应保存到缓存的条件，如果字符串参数的至少一个值不为空且不等于"0"，则将不保存该响应到缓存
```
语法：proxy_no_cache string;
上下文：http、server、location
示例：proxy_no_cache $http_pragma    $http_authorization;
```

## proxy_cache_bypass 
>- 定义条件，在该条件下将不会从缓存中获取响应
```
语法：proxy_cache_bypass string;
上下文：http、server、location
示例：proxy_cache_bypass $http_pragma    $http_authorization;
```

## upstream_cache_status 变量
>- 它存储了缓存是否命中的信息，会设置在响应头信息中，在调试中非常有用
>- MISS: 未命中缓存
>- HIT: 命中缓存
>- EXPIRED: 缓存过期
>- STALE: 命中了陈旧缓存
>- REVALIDDATED: nginx 验证陈旧缓存依然有效
>- UPDATEING: 内容陈旧，但正在更新
>- BYPASS: 响应从原始服务器获取

## docker 配置
>- 准备文件
```
mkdir -p /data/nginx/{proxy_cache,upstream1,upstream2}/html

cat >/data/nginx/proxy_cache/html/index.html<<-EOF
<h1>cache proxy html</h1>
EOF

cat >/data/nginx/upstream1/html/index.html<<-EOF
<h1>upstream1 html</h1>
EOF

cat >/data/nginx/upstream2/html/index.html<<-EOF
<h1>upstream2 html</h1>
EOF

cat | tee /data/nginx/{upstream1,upstream2}/default.conf<<-EOF
server {
 listen 80;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
}
EOF

cat >/data/nginx/proxy_cache/default.conf<<-EOF
proxy_cache_path /etc/nginx/cache_temp levels=2:2 keys_zone=cache_zone:30m max_size=2g inactive=60m use_temp_path=off;

upstream backend_server{
 server 191.168.126.71;
 server 191.168.126.72;
}

server {
 listen 80;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
 
 ## URI 中后缀为 .txt|.text 的设置变量值 
 if (\$request_uri ~ \\.(txt|text)\$){
   set \$cache_name "no cache";
 }
 location /proxy_cache/ {
  proxy_no_cache \$cache_name; ## 判断该变量是否有值，如果有值则不进行缓存，如果没有值则进行缓存
  proxy_cache cache_zone; ## 设置缓存内存，上面配置中已经定义好的
  proxy_cache_valid 200 5m; ## 缓存状态为200的请求，缓存时长为5分钟
  proxy_cache_key \$request_uri; ## 缓存文件的key为请求的URI
  add_header Nginx-Cache-Status \$upstream_cache_status; # 把缓存状态设置为头部信息，响应给客户端
  proxy_pass http://backend_server/;
 }
}
EOF

```

>- 自定义网络
```
docker network create --subnet=191.168.126.0/24 mynet
docker network ls 
docker network inspect mynet 
```

>- 运行

```
docker run --network=mynet --name=nginx_upstream1 --ip=191.168.126.71 -d \
-v /data/nginx/upstream1/html:/usr/share/nginx/html \
-v /data/nginx/upstream1/default.conf:/etc/nginx/conf.d/default.conf \
-p 8001:80 nginx:1.20.0 

docker run --network=mynet --name=nginx_upstream2 --ip=191.168.126.72 -d \
-v /data/nginx/upstream2/html:/usr/share/nginx/html \
-v /data/nginx/upstream2/default.conf:/etc/nginx/conf.d/default.conf \
-p 8002:80 nginx:1.20.0 


docker run --rm --network=mynet --name=nginx_proxy_cache \
-v /data/nginx/proxy_cache/html:/usr/share/nginx/html \
-v /data/nginx/proxy_cache/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:80 nginx:1.20.0 nginx -t

docker run --rm --network=mynet --name=nginx_proxy_cache \
-v /data/nginx/proxy_cache/html:/usr/share/nginx/html \
-v /data/nginx/proxy_cache/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:80 nginx:1.20.0 nginx -g "daemon off;"

docker run --network=mynet --name=nginx_proxy_cache -d \
-v /data/nginx/proxy_cache/html:/usr/share/nginx/html \
-v /data/nginx/proxy_cache/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:80 nginx:1.20.0

```

>- 检查配置
```
docker inspect nginx_upstream1
docker inspect nginx_upstream2

docker run --rm -it --name=busybox --network=mynet busybox
ping nginx_upstream1
ping nginx_upstream2
ping 191.168.126.71
ping 191.168.126.72

docker exec nginx_upstream1 nginx -t
docker exec nginx_upstream2 nginx -t
docker exec nginx_proxy_cache nginx -t
docker exec nginx_proxy_cache nginx -T
```

>- 测试
```
curl localhost:8001
curl localhost:8002
curl localhost:8003

curl localhost:8000
curl -I localhost:8000/proxy/
curl -i localhost:8000/proxy/

docker cp nginx_proxy_cache:/etc/nginx/cache_temp ./

```

>- 清理
```
docker rm -f nginx_upstream1 nginx_upstream2 nginx_proxy_cache
docker network rm mynet
```