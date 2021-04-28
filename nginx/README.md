## nginx
>- http://nginx.org/

## CentOS7 下安装
>- http://nginx.org/en/linux_packages.html#RHEL-CentOS

```
cat >/etc/yum.repos.d/nginx.repo<<EOF
[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true
EOF

yum install -y nginx

nginx -v
```

## 源码安装
>- 编译环境  yum install -y gcc 
>- 正则表达式库 yum install -y pcre pcre-devel
>- 解压缩库  yum install -y zlib zlib-devel
>- 安全套接字密码库及 ssl协议 yum install -y openssl openssl-devel
>- yum install -y gcc pcre pcre-devel zlib zlib-devel openssl openssl-devel

```
wget http://nginx.org/download/nginx-1.20.0.tar.gz
tar -xf nginx-1.20.0.tar.gz
cd nginx-1.20.0
./configure 
make 
make install
```

## 命令
>- 启动 nginx 
>- 停止 nginx -s stop|quit
>- 重新加载 nginx -s reload  

## Docker 启动
>- docker pull nginx:1.20.0
>- docker run --name=nginx -d -p 8080:80 nginx:1.20.0
>- docker run --rm --name=nginx -p 8080:80 nginx:1.20.0 nginx -g "daemon off;"

## Docker-compose 启动
```
cat >docker-compose.yml<<EOF 
web:
 image: nginx:1.20.0
 ports:
  - 80:80 
 command: [nginx,'-g','daemon off;']
EOF
```

>- 启动 docker-compose up
>- 停止 docker-compose down