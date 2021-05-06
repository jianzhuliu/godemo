## sysctl 
>- /usr/sbin/sysctl

## 参数类型
```
sysctl -a | awk -F '.' '{print $1}' | sort -k1 | uniq
```

## 网络相关
```
sysctl -a |grep '^net.' | awk -F '[.| ]' '{print $2}' | sort -k1 | uniq
```

## 缓存区配置
```
sysctl -a | grep '^net.' | grep '[r|w|_]mem[_| ]'
```

>- mem、rmem、wmem 分别是总内存、接收缓冲区内存、发送缓冲区内存
>- rmem 和 wmem 的单位都是“字节”，而 mem 的单位是“页”
>- “页”是操作系统管理内存的最小单位，在 Linux 系统里，默认一页是 4KB 大小
>- tcp_mem、tcp_rmem、tcp_wmem、udp_mem 这几个参数后面有三个值

>- 对于 tcp_rmem 和 tcp_wmem 来说，这三个值是单个套接字可分配内存的大小，从左到右分别是最小值、默认值、最大值
>- 这当中的默认值和最大值会分别被 net.core 下对应的 default 值和 max 值覆盖

>- 对于 tcp_mem 和 udp_mem 来说，它后面的三个值用于控制内存压力，从左到右分别是内存压力的最小值、压力值、最大值
>- 比如 tcp_mem 的最小值是 188964、压力值是251954、最大值是 377928。当 TCP 总内存使用量小于 188964 时，表示内存毫无压力，不用考虑回收；当内存使用量超过 251954 时，系统会开始回收内存，直到小于 188964；当内存使用量达到 377928 时，系统将会拒绝分配套接字，并输出日志“TCP: too many of orphaned sockets”。

>- 对于秒杀接口这种大量短连接的业务场景，需要减少 rmem 和 wmem 相关的数值。比如将最小值、默认值、最大值分别改为 4096、4096、8192，就能建立更多的连接

## TCP 协议参数
>- sysctl -a | grep '^net.ipv4.tcp_'

```
# 抵御攻击者用大量 SYN 报文发起的短报文攻击
net.ipv4.tcp_syncookies = 1
# 避免网络抖动后重传已成功发送的数据包,设置选择性重传的参数
net.ipv4.tcp_sack = 1
# 以免文件描述符不够用导致性能问题
fs.file-max = 65535


# 重用处于 TIME-WAIT 状态的套接字
net.ipv4.tcp_tw_reuse = 1

# 快速回收 TIME-WAIT 状态的套接字
net.ipv4.tcp_tw_recycle = 1

# 关闭处于 FIN-WAIT-2 状态 30 秒以上的套接字
net.ipv4.tcp_fin_timeout = 30

# 设置空闲 TCP 连接存活时间，以便即时关闭空闲连接，回收资源
net.ipv4.tcp_keepalive_time=1800


```
