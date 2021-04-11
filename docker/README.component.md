## Docker 组件(component)

## docker
>- Docker 客户端一个完整实现，负责发送 Docker 操作请求

## dockerd
>- Docker 服务端的后台常驻进程，负责接收客户端请求并返回请求结果

## docker-init
>- 当业务主进程没有进程回收能力时，docker-init 可以作为容器的1号进程，负责管理容器内子进程
>- docker run -it busybox sh 
>- ps aux -- 1号进程就是 sh 
>- docker run -it --init busybox sh 
>- ps aux -- 1号进程就是 /sbin/docker-init -- sh

## docker-proxy 
>- 用来做 Docker 的网络实现，通过设置 iptables 规则使得访问到主机的流量可以被顺利转发到容器中
>- docker run --name=nginx -d -p 8080:80 nginx
>- docker inspect nginx | grep -i IPAddress
>- ps aux | grep docker-proxy
>- iptables -L -nv -t nat 
>- curl http://localhost:8080

## containerd 
>- 负责容器生命周期的管理
>- 镜像的管理
>- 接收 dockerd 的请求，通过适当的参数调用 runc 启动容器
>- 管理存储相关资源
>- 管理网络相关资源

## containerd-shim
>- 将 containerd 和真正的容器进程解耦,使用 containerd-shim 作为容器进程的父进程，可以实现重启 containerd 不影响已经启动的容器进程

## runc 
>- 容器运行时组件，通过调用 Namespace、cgroups 等系统接口，实现容器的创建和销毁
>- mkdir /tmp/runc && cd /tmp/runc
>- mkdir rootfs && docker export $(docker create busybox) | tar -C rootfs -xvf -
>- runc spec
>- runc run busybox
>- 另外开启 shell 窗口， runc list
>- runc kill -a busybox
>- runc delete busybox
>- rm -rf /tmp/runc

## 弊端
>- 调用链过长，出现问题需要分析很多组件才能定位问题


