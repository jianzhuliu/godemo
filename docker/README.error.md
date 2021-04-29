## 问题汇总

## WARNING: IPv4 forwarding is disabled. Networking will not work.
>- 修改文件 /usr/lib/sysctl.d/00-system.conf
>- 添加  net.ipv4.ip_forward=1
>- 重启网卡 service network restart
>- 检查配置是否生效 sysctl net.ipv4.ip_forward

