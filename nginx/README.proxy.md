## 反向代理

## docker 自定义网络类型
>- 文件准备
```
mkdir -p /data/nginx/{upstream,proxy}/html

cat >/data/nginx/upstream/html/index.html<<-EOF
<h1>8001:upstream html</h1>
EOF

cat >/data/nginx/proxy/html/index.html<<-EOF
<h1>8000:proxy html</h1>
EOF

cat >/data/nginx/upstream/default.conf<<-EOF
server {
 listen 8001;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
}
EOF

cat >/data/nginx/proxy/default.conf<<-EOF
upstream backend_server{
 server nginx_upstream:8001 weight=3 max_conns=1000 fail_timeout=10s max_fails=2;
 keepalive 32;
 keepalive_requests 50;
 keepalive_timeout 30s;
}

server {
 listen 8000;
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

>- 新建网络
```
docker network create nginx 
docker network ls 
docker network inspect nginx 
```

>- 后端运行
```
docker run --name=nginx_upstream -d \
-v /data/nginx/upstream/html:/usr/share/nginx/html \
-v /data/nginx/upstream/default.conf:/etc/nginx/conf.d/default.conf \
-p 8001:8001 nginx:1.20.0

docker run --name=nginx_proxy -d \
-v /data/nginx/proxy/html:/usr/share/nginx/html \
-v /data/nginx/proxy/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:8000 nginx:1.20.0
```

>- 或者前端运行
```
docker run --network=nginx --rm --name=nginx_upstream \
-v /data/nginx/upstream/html:/usr/share/nginx/html \
-v /data/nginx/upstream/default.conf:/etc/nginx/conf.d/default.conf \
-p 8001:8001 nginx:1.20.0 nginx -g "daemon off;"

docker run --network=nginx --rm --name=nginx_proxy \
-v /data/nginx/proxy/html:/usr/share/nginx/html \
-v /data/nginx/proxy/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:8000 nginx:1.20.0 nginx -g "daemon off;"
```

>- 测试
```
curl localhost:8000
curl localhost:8000/proxy/
curl localhost:8001

```

>- 清理
```
docker rm -f nginx_upstream nginx_proxy
docker network rm nginx
```

>- 上游服务器指定ip
```
docker network create --subnet=191.168.126.0/24 mynet
docker network inspect mynet 
cat >/data/nginx/proxy/default.conf<<-EOF
upstream backend_server{
 server 191.168.126.71:8001 weight=3 max_conns=1000 fail_timeout=10s max_fails=2;
 keepalive 32;
 keepalive_requests 50;
 keepalive_timeout 30s;
}

server {
 listen 8000;
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

docker run --network=mynet --rm --name=nginx_upstream --ip=191.168.126.71 \
-v /data/nginx/upstream/html:/usr/share/nginx/html \
-v /data/nginx/upstream/default.conf:/etc/nginx/conf.d/default.conf \
-p 8001:8001 nginx:1.20.0 nginx -g "daemon off;"

docker run --network=mynet --rm --name=nginx_proxy \
-v /data/nginx/proxy/html:/usr/share/nginx/html \
-v /data/nginx/proxy/default.conf:/etc/nginx/conf.d/default.conf \
-p 8000:8000 nginx:1.20.0 nginx -g "daemon off;"

curl localhost:8000
curl localhost:8000/proxy/
curl localhost:8001

docker network rm mynet 

```

>- 访问 localhost:8000/proxy/ 通过 upstream 找到 191.168.126.71:8001/