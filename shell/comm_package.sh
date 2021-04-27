#/bin/bash
### 常用linux工具包

## 查看命令帮助文档
yum -y install man

## 追踪二进制执行文件系统调用
yum -y install strace
## strace -ff -o out command

## 模拟网络通讯
yum -y install nc 
## nc localhost 8001

## 查看目录下树形结构
yum -y install tree 

## tcp协议日志监听，三次握手，四次分手日志记录可见
yum -y install tcpdump
## tcpdump -nn -i eth0 port 80

## 网络文件下载
yum -y install wget

## make
yum -y install make 

## lsof 展示被进程打开的文件信息
yum -y install lsof
## lsof -p 10765

## 格式转换
yum -y install dos2unix
