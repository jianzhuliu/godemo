## Network 网络模式

#### null 空网络模式
>- 处理一些保密数据，出于安全考虑
>- docker run --net=none -it busybox
>- ifconfig
>- route -n

#### bridge 桥接模式
>- 容器启动时，默认的网络模式
>- Linux veth -- 虚拟设备接口，成对出现，充当一个桥梁，连接虚拟网络设备
>- Linux bridge -- 虚拟设备，用来连接网络的设备
>- Docker 启动时，libnetwork会在主机上创建 docker0 网桥，实现网络互通

#### host 主机网络模式
>- 容器需要控制主机网络或者用主机网络提供服务
>- docker run --net=host -it busybox sh
>- ifconfig
>- route -n
>- ip a

#### container 网络模式
>- 两个容器之间需要之间通过 localhost 通信，一般用于网络接管或者代理服务
>- docker run -d --name=busybox1 busybox sleep 3600
>- docker exec -it busybox1 sh 
>- ifconfig 
>- docker run -it --net=container:busybox1 --name=busybox2 busybox sh 
>- ifconfig