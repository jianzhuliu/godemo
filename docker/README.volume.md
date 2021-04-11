## Volume 数据卷

#### 创建数据卷
>- docker volume create myvolume

#### 查看数据卷
>- docker volume ls   --列表
>- docker volume inspect myvolume -- 查看详细信息

#### 使用数据卷
>- docker run -d -p 8080:80 --name=nginx --mount source=myvolume,target=/usr/share/nginx/html nginx
>- curl localhost:8080
>- docker exec -it nginx sh   --进入容器

```
## 使用以下内容直接替换 /usr/share/nginx/html/index.html 文件 

cat <<EOF >/usr/share/nginx/html/index.html
<!DOCTYPE html>
<html>
<head>
<title>Hello, Docker Volume!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Hello, Docker Volume!</h1>
</body>
</html>
EOF

```

>- curl localhost:8080
>- docker rm -f nginx
>- docker run -d -p 8080:80 --name=nginx --mount source=myvolume,target=/usr/share/nginx/html nginx
>- curl localhost:8080
>- docker exec -it nginx sh 
>- cat /usr/share/nginx/html/index.html
>- docker rm -f nginx
>- cat /var/lib/docker/volumes/myvolume/_data/index.html
>- docker volume rm myvolume
>- ls -l /var/lib/docker/volumes/

#### 删除数据卷
>- docker volume rm myvolume

#### 容器之间数据共享
>- docker volume create mylog
>- ls -l /var/lib/docker/volumes/mylog/_data/
>- docker run --mount source=mylog,target=/tmp/log --name=producer -it busybox

```
cat >/tmp/log/mylog.log<<EOF
Hello, My log.
EOF 
```

>- cat /var/lib/docker/volumes/mylog/_data/mylog.log
>- 新开一个窗口， docker run -it --name=consumer --volumes-from producer busybox
>- docker rm -f producer consumer
>- docker volume rm mylog

#### 主机与容器之间数据共享
>- docker run -it -p 8080:80 --name=nginx -v /data/nginx/www:/usr/share/nginx/html nginx

```
cat <<EOF > /data/nginx/www/index.html
<html>
<head><title>Hello</title></head>
<body>
<center><h1>World</h1></center>
</body>
</html>
EOF

cat /data/nginx/www/index.html
```

>- curl localhost:8080
>- docker rm -f nginx

## Docker 卷的实现原理
>- Docker 容器的文件系统不是一个真正的文件系统，而是通过联合文件系统实现的一个伪文件系统
>- Docker 卷则是直接利用主机的某个文件或者目录，它可以绕过联合文件系统，直接挂载主机上的文件或目录到容器中
>- docker volume create myvolume
>- ls -l /var/lib/docker/volumes/myvolume/_data/
>- docker run -it --name=busybox --mount source=myvolume,target=/data busybox 
>- echo "hello" > /data/mylog.log
>- ls -l /var/lib/docker/volumes/myvolume/_data/
>- cat /var/lib/docker/volumes/myvolume/_data/mylog.log
>- docker rm -f busybox && docker volume rm myvolume
>- ls -l /var/lib/docker/volumes/
>- Docker 卷的实现原理是在主机的 /var/lib/docker/volumes 目录下，根据卷的名称创建相应的目录，然后在每个卷的目录下创建 _data 目录，在容器启动时如果使用 --mount 参数，Docker 会把主机上的目录直接映射到容器的指定目录下，实现数据持久化
