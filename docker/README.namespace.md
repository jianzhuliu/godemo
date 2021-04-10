## 资源隔离 (Namespace)
Namespace 是 Linux 内核的一个特性，该特性可以实现在同一主机系统中，对进程 ID、主机名、用户 ID、文件名、网络和进程间通信等资源的隔离。

## Mount Namespace -- 隔离挂载点
>- unshare --mout --fork /bin/bash
>- mkdir /tmp/tmpfs
>- mount -t tmpfs -o size=20m tmpfs /tmp/tmpfs
>- df -h  
>- 另外开启一个窗口执行 df -h 对比可知
>- 同理对比 ls -s /proc/self/ns 只有 mnt对应的id不同
>- umount /tmp/tmpfs

## PID Namespace -- 隔离进程ID
>- ps aux
>- unshare --pid --fork --mount-proc /bin/bash
>- ps aux  将看不到主机进程

## UTS Namespace -- 隔离主机名和域名
>- hostname
>- unshare --uts --fork /bin/bash
>- hostname -b newhostname
>- hostname

## IPC Namespace -- 隔离进程间通信
>- unshare --ipc --fork /bin/bash
>- ipcs -q  --查看当前 IPC Namespace 下的系统通信队列列表
>- ipcmk -Q --创建一个系统通信
>- ipcs -q

## User Namespace -- 隔离用户及用户组
>- 使用非 root 用户登录
>- unshare --user -r /bin/bash
报错 unshare: unshare failed: Invalid argument

>- root用户操作 cat /proc/sys/user/max_user_namespaces
>- root用户操作 echo 65535 > /proc/sys/user/max_user_namespaces

>- unshare --user -r /bin/bash
>- id
>- reboot   -- 并不能获取到主机root权限

## NET Namespace -- 隔离网络设备、IP地址和端口
>- ip a 
>- unshare --net --fork /bin/bash


