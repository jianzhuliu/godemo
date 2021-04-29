## Network 网络模式
>- docker network ls 

#### null 空网络模式
>- 处理一些保密数据，出于安全考虑
>- docker run --net=none -it busybox
>- ifconfig
>- route -n

#### bridge 桥接模式
>- 容器启动时，默认的网络模式, 虚拟接口
>- Linux veth -- 虚拟设备接口，成对出现，充当一个桥梁，连接虚拟网络设备
>- Linux bridge -- 虚拟设备，用来连接网络的设备
>- Docker 启动时，libnetwork会在主机上创建 docker0 网桥，实现网络互通
>- 虚拟接口的优势就是转发效率极高（因为Linux是在内核中进行数据的复制来实现虚拟接口之间的数据转发，无需通过外部的网络设备交换）

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
>- docker run -it --rm --net=container:busybox1 --name=busybox2 busybox sh 
>- ifconfig

## docker0
>- 默认 bridge 会创建一个 docker0 网桥, 网关地址为 172.17.0.1 的路由器
>- 随机分配 IP 地址，IP 地址范围为 172.17.0.2—172.17.255.254

## 172.17.0.1/16 
>- IP 后面跟着一个斜杠+数字，比如/16 这种的，表示子网掩码位数，可以区分 IP 段类型/16 就是子网有 16 个一（二进制），剩余全为零
>- 而子网一共分为 4 段，每段 8 位，所以二进制写出来也就是： 11111111.11111111.00000000.00000000
>- 而 8 个二进制的 1 也就等于十进制的 255，转换成十进制也就是：255.255.0.0
>- 整个 172.17.0.0/16 的 IP 包含范围是 172.17.0.0~172.17.255.255，去掉首尾两个特殊地址，实际可分配 IP 范围为 172.17.0.1~172.17.255.254

#### 自定义网络
>- 创建一个名为 mynet，默认类型为 bridge, 子网掩码为 255.255.255.0,可用网络段为 171.168.1.1 ~ 171.168.1.254

```
docker network create --subnet=171.168.1.0/24 mynet 
docker network ls 
docker network inspect mynet
```

>- 运行容器，指定ip 
```
docker run --rm -it --name=busybox1 --network=mynet --ip=171.168.1.16 busybox
ifconfig
```

>- 其他容器访问
```
docker run --rm -it --name=busybox2 --network=mynet busybox
ping busybox1
ping 171.168.1.16
ifconfig

```

>- 删除网络
```
docker network rm mynet
```
