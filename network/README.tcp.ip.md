## 网络协议  tcp/ip

## TCP/IP 四层协议
>- 应用层 	http|ftp|pop3
>- 传输控制层 tcp (Transmission Control Protocol|传输控制协议)
>- 网络层 
```
ip(Internet Protocol) 网际协议
icmp(Internet Control Message Protocol|Internet) 控制报文协议
arp(Address Resolution Protocol) 地址解析协议
```

>- 数据链路层

## 命令 exec 
>- shell 内建命令将不启动新的shell,而是用被执行命令替换当前的shell进程，老进程的环境被清理叼
>- exec command;  command 替代shell程序，命令退出，shell退出，如 exec ls 
>- exec 文件重定向，比如 exec 5<file; 关闭文件描述符 exec 5<&-

## I/O 重定向
>- fd 文件描述符 file descriptor
>- stdin,标准输入(0)
>- stdout,标准输出(1)
>- stderr,标准错误输出(2)

## man 帮助命令
>- yum install -y man man-pages
>- man 1 : shell 命令,如ls,vim... man 1 open
>- man 2 : 系统调用，open,close,accpet,socket... man 2 open 
>- man 3 : 库函数，printf,fopen 
>- man 4 : /dev下的文件
>- man 5 : 一些配置文件的格式
>- man 6 : 预留
>- man 7 : 附件和变量
>- man 8 : 只能由 root 执行的系统管理命令，如mkfs 

## socket 套接字, ip:port -> ip:port 成对出现
>- 创建到百度80端口的文件描述符9
```
exec 9<> /dev/tcp/www.baidu.com/80
```

>- 构造请求头, 标准输出1 重定向到文件描述符9
```
echo -e "GET / HTTP/1.1\n" 1>& 9
```

>- 读取数据, 标准输入0 数据来源于文件描述符9
```
cat 0<& 9
```

>- 查看当前进程id 
```
echo $$ 
```

>- 查看当前进程下所有文件描述符列表
```
ls -l /proc/$$/fd
```

>- 关闭文件描述符 
```
exec 9<&-
```

## TCP 面向连接、可靠的协议
>- 三次握手,互相能建立连接，分配资源
```
c --->sync     	---> s
s --->sync ack 	---> c
c --->act   	---> s 
```

>- 四次挥手,互相能正确关闭资源
```
c --->fin     	---> s
s --->ack     	---> c
s --->fin     	---> c
c --->ack     	---> s
```

>- 抓包验证 
```
yum install -y tcpdump

###在 ens33 网卡上监听80端口
tcpdump -nn -i ens33 port 80
curl www.baidu.com

-nn Don't convert protocol and port numbers etc. to names either.(也不要将协议和端口号等转换为名称。)
-i interface 网关
```

## 路由表
>- route -n
>- ping www.baidu.com (183.232.231.174)
>- 183.232.231.174 与 Genmask子网掩码and运算，匹配成功，则与 Destination比较，如果成功就走下一跳 Gateway
>- 上步失败后，就与 0.0.0.0 匹配，走下一跳 Gateway

## arp (Address Resolution Protocol) 地址解析协议
>- arp -n 
>- 在 ens33 网卡上监听80端口或者 arp 协议
```
tcpdump -nn -i ens33 port 80 or arp 
arp -d 192.168.126.1 && curl www.baidu.com
```

