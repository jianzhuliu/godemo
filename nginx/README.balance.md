## 负载均衡

## docker 自定义网络
>- 准备文件
```
mkdir -p /data/nginx/{proxy,upstream1,upstream2,upstream3}/html

cat >/data/nginx/proxy/html/index.html<<-EOF
<h1>proxy html</h1>
EOF

cat >/data/nginx/upstream1/html/index.html<<-EOF
<h1>upstream1 html</h1>
EOF

cat >/data/nginx/upstream2/html/index.html<<-EOF
<h1>upstream2 html</h1>
EOF

cat >/data/nginx/upstream3/html/index.html<<-EOF
<h1>upstream3 html</h1>
EOF

cat | tee /data/nginx/{upstream1,upstream2,upstream3}/default.conf<<-EOF
server {
 listen 80;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
}
EOF

cat >/data/nginx/proxy/default.conf<<-EOF
upstream backend_server{
 server 191.168.126.71;
 server 191.168.126.72;
 server 191.168.126.73;
}

server {
 listen 80;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
 location /proxy/ {
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

docker run --network=mynet --name=nginx_upstream3 --ip=191.168.126.73 -d \
-v /data/nginx/upstream3/html:/usr/share/nginx/html \
-v /data/nginx/upstream3/default.conf:/etc/nginx/conf.d/default.conf \
-p 8003:80 nginx:1.20.0 

docker run --network=mynet --name=nginx_proxy -d \
-v /data/nginx/proxy/html:/usr/share/nginx/html \
-v /data/nginx/proxy/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:80 nginx:1.20.0

```

>- 检查配置
```
docker inspect nginx_upstream1
docker inspect nginx_upstream2
docker inspect nginx_upstream3

docker run --rm -it --name=busybox --network=mynet busybox
ping nginx_upstream1
ping nginx_upstream2
ping nginx_upstream3
ping 191.168.126.71
ping 191.168.126.72
ping 191.168.126.73

docker exec nginx_upstream1 nginx -t
docker exec nginx_upstream2 nginx -t
docker exec nginx_upstream3 nginx -t
docker exec nginx_proxy nginx -t
docker exec nginx_proxy nginx -T
```

>- 测试
```
curl localhost:8001
curl localhost:8002
curl localhost:8003

curl localhost:8000
curl localhost:8000/proxy/

```

>- hash 算法 (hash $request_uri;)
``` 
cat >/data/nginx/proxy/default.conf<<-EOF
upstream backend_server{
 hash \$request_uri;
 server 191.168.126.71;
 server 191.168.126.72;
 server 191.168.126.73;
}

server {
 listen 80;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
 location /proxy/ {
  proxy_pass http://backend_server/;
 }
}
EOF

docker exec nginx_proxy nginx -t
docker exec nginx_proxy nginx -s reload
docker exec nginx_proxy nginx -T

docker stop nginx_proxy
docker start nginx_proxy

curl localhost:8000/proxy/
curl localhost:8000/proxy/?a=b
curl localhost:8000/proxy/?a=c
curl localhost:8000/proxy/?a=d

```

>- ip_hash (ip_hash;)
>- 最少连接数 (least_conn;)
```
zone myproxy 10M; ##zone 可以设置共享内存空间名称和大小 
least_conn; ##最少连接数
```


>- 清理
```
docker rm -f nginx_upstream1 nginx_upstream2 nginx_upstream3 nginx_proxy
docker network rm mynet
```