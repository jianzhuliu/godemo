# docker
https://www.docker.com/

## 文档
https://docs.docker.com/get-started/

## 官方仓库
https://hub.docker.com/
https://labs.play-with-docker.com/

## 镜像
[image](./README.image.md)

## 容器
[container](./README.container.md)

## 仓库
[container](./README.repository.md)

## 对比
> 镜像包含了容器运行所需要的文件系统结构和内容，是静态的只读文件，
> 而容器则是在镜像的只读层上创建了可写层，并且容器中的进程属于运行状态，容器是真正的应用载体。
> 仓库是存储与分发镜像的地方

## 监控 cadvisor
>- from https://github.com/google/cadvisor/blob/master/docs/running.md
```
docker run \
--volume=/:/rootfs:ro \
--volume=/var/run:/var/run:rw \
--volume=/sys/fs/cgroup/cpu,cpuacct:/sys/fs/cgroup/cpuacct,cpu \
--volume=/var/lib/docker/:/var/lib/docker:ro \
--publish=8080:8080 \
--detach=true \
--name=cadvisor \
--privileged=true \
google/cadvisor:latest

```

## 容器与虚拟机
>- 虚拟机是通过管理系统(Hypervisor)模拟出CPU、内存、网络等资源，创建客户内核和操作系统
>- 虚拟机有自己的内核和操作系统，不会直接使用到主机的操作系统和硬件资源，隔离性鸡安全性有更好的保障

>- Docker容器是通过Linux内核的Namespace技术实现文件系统、进程、设备及网络的隔离
>- 通过Cgroups对CPU、内存等资源进行限制。隔离性靠内核提供
>- 容器的性能消耗非常小，镜像也非常小，秒级启动

## 安全
>- Docker 自身安全 	-- CVE收录,权限提升、信息泄露等
>- 镜像软件、仓库、用户程序漏洞 -- 容器基于镜像创建并启动
>- Linux内核隔离性不够 	-- 部分关键内容没有被安全隔离，比如/sys、 /proc等
>- 所有容器共用主机内核

>- 使用最新的Docker版本
>- 使用受信任的镜像仓库，且通过https协议通信，并安装镜像安全扫描组件
>- 及时升级内核补丁，使用 Capabilities 划分权限，使用安全加固组件
>- 资源限制, docker run --cpus=1 -m=2048m --pids-limit=1000 
>- 使用安全容器，如 Kata Containers

## 资源隔离
[Namespace](./README.namespace.md)

## Cgroups -- control groups
[Cgroups](./README.cgroups.md)