## 容器 (Container)

>- 容器是基于镜像创建的可运行实例，并且单独存在，一个镜像可以创建出多个容器

## 容器的生命周期

>- created: 初建状态 -- docker create 
>- running: 运行状态 -- docker run 
>- stopped: 停止状态 -- docker stop 
>- paused: 	暂停状态 -- docker pause 
>- deleted: 删除状态 -- docker rm

## 创建并启动容器
>- docker create -it --name=busybox busybox
>- docker run -it --name=busybox busybox
>- 参数 -t 分配一个伪终端, -i 支持终端的stdin, 同时使用就可以进入交互模式

#### docker run 创建并启动容器时，docker后台执行流程
>- Docker会检查本地是否存在 busybox 镜像，如果不存在，就去docker hub 拉取 busybox 镜像
>- 使用 busybox 镜像创建并启动一个容器
>- 分配文件系统，并且在镜像只读层外创建一个读写层
>- 从Docker ip池中分配一个ip给容器
>- 执行用户的启动命令，运行镜像

## 终止容器
>- docker stop busybox
>- 查看已经停止状态的容器 docker ps -a
>- 首先会向运行中的容器发送 SIGTERM 信号，如果容器内进程接受并能够处理 SIGTERM，则等待处理完毕后退出，如果等待一段时间后，容器仍然没有退出，则会发送 SIGKILL 强制终止容器

## 进入容器
>- docker attach busybox
>- docker exec -it busybox sh (单独开启一个进程，相互独立互不干扰)
 
## 删除容器
>- docker rm busybox

## 导出容器
>- docker export busybox > busybox.tar

## 导入容器
>- docker import busybox.tar busybox:test




