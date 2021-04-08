## 镜像
镜像是一个只读的 Docker 容器模板，包含启动容器所需要的所有文件系统结构和内容。
镜像是一个特殊的文件系统，它提供了容器运行时所需的程序、软件库、资源、配置等静态数据。
镜像不包含任何动态数据，镜像内容在构建后不会被改变。
Docker 镜像是静态的分层管理的文件组合，镜像底层的实现依赖于联合文件系统（UnionFS）

#### 拉取镜像 docker pull [Registry]/[Repository]/[Image]:[Tag]
>- Registry 注册服务器,默认 docker.io
>- Repository 镜像仓库,默认 library
>- Image 镜像名称
>- Tag 镜像标签，默认 latest

docker pull busybox
docker pull docker.io/library/busybox:latest

#### 查看镜像 
>- docker images 
>- docker image ls 

#### 镜像重命名
>- docker tag busybox:latest mybusybox:latest 

#### 删除镜像
>- docker rmi mybusybox 
>- docker image rm mybusybox


#### 构建镜像
###### docker commit 从运行中的容器提交为镜像
> 1 开启 busybox 容器
>- docker run --name=busybox --rm -it busybox sh 
>- touch hello.txt | echo "Hello World" > hello.txt 

> 2 新开窗口
docker commit busybox busybox:hello 

> 3 导出容器内镜像文件
>- mkdir -p /tmp/busybox && cd /tmp/busybox
>- docker export $(docker create busybox:hello) -o /tmp/busybox/busybox.tar
>- tar -xf busybox.tar

###### docker build 从 Dockerfile 构建镜像
>- FROM 	Dockerfile 除了注释第一行必须是 FROM ，FROM 后面跟镜像名称，代表我们要基于哪个基础镜像构建我们的容器。
>- RUN 		RUN 后面跟一个具体的命令，类似于 Linux 命令行执行命令
>- ADD 		拷贝本机文件或者远程文件到镜像内
>- COPY 	拷贝本机文件到镜像内
>- USER 	指定容器启动的用户
>- ENTRYPOINT	容器的启动命令
>- CMD 		CMD 为 ENTRYPOINT 指令提供默认参数，也可以单独使用 CMD 指定容器启动参数
>- ENV		指定容器运行时的环境变量，格式为 key=value
>- ARG		定义外部变量，构建镜像时可以使用 --build-arg = 的格式传递参数用于构建
>- EXPOSE	指定容器监听的端口，格式为 [port]/tcp 或者 [port]/udp
>- WORKDIR 	为 Dockerfile 中跟在其后的所有 RUN、CMD、ENTRYPOINT、COPY 和 ADD 命令设置工作目录。

> demo 
```
mkdir -p /tmp/mynginx && cd /tmp/mynginx
cat >Dockerfile<<-EOF
FROM centos:7
COPY nginx.repo /etc/yum.repos.d/nginx.repo
RUN yum install -y nginx
EXPOSE 80
ENV HOST=mynginx
CMD ["nginx","-g","daemon off;"]
EOF

cat >nginx.repo<<-EOF
#### from http://nginx.org/en/linux_packages.html#RHEL-CentOS

[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/\$releasever/\$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true

[nginx-mainline]
name=nginx mainline repo
baseurl=http://nginx.org/packages/mainline/centos/\$releasever/\$basearch/
gpgcheck=1
enabled=0
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true
EOF

docker build -t mynginx .

docker run -p 8090:80 mynginx

curl 192.168.195.71:8090
```

## 镜像实现原理
镜像是由一系列的镜像层（layer ）组成，每一层代表了镜像构建过程中的一次提交，当我们需要修改镜像内的某个文件时，只需要在当前镜像层的基础上新建一个镜像层，并且只存放修改过的文件内容。分层结构使得镜像间共享镜像层变得非常简单和方便

