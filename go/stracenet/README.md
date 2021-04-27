## 跟踪 net.Listen 系统调用
>- 安装命令工具
>- yum install -y strace nc tcpdump lsof
>- netstat 属于基本网络工具包 yum install -y net-tools.x86_64

## 准备测试文件， main.go , 并编译 go build -o server main.go

## 第1个窗口
>- strace -ff -o out ./server
>- -ff 监听所有进程及其创建出来的所有子进程
>- -o 输出文件名，多进程，文件名加.进程号
>- 输出结果可见监听端口8001及pid,比如 pid 10765

## 第2个窗口
>- 当前目录下，可见out开头的文件
>- 该进程下所有的子进程 ll /proc/10765/task
>- 该进程下所有的文件描述符 ll /proc/10765/fd
>- lsof -p 10765  
>- netstat -natp 
>- tail -f out.10765 

## 第3个窗口
>- 监听 tcp 三次握手及四次分手
>- tcpdump --nn -i lo port 8001
>- --nn ip使用数字显示
>- -i 在哪个网卡上监听, lo 回环网卡
>- port 8001 监听端口8001 

## 第4个窗口
>- 查看连接状态
>- losf -p 10765
>- netstat -natp | grep server 

## 第5个窗口
>- 模拟建立tcp连接，并交互操作
>- nc localhost 8001
