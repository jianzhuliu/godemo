#/bin/bash

#设置硬件时钟与本地时钟一致
timedatectl set-local-rtc 1

#设置时区
timedatectl set-timezone Asia/Shanghai

#安装 ntpdate 
# Network Time Protocol
yum -y install ntpdate

# 同步时间
# pool.ntp.org
# cn.ntp.org.cn
# hk.ntp.org.cn
# us.ntp.org.cn


ntpdate -u cn.ntp.org.cn

date

# 一般使用 crontab 来同步时间
# yum -y install crontab
# crontab -e
# */5 * * * * /usr/sbin/ntpdate cn.ntp.org.cn
# service crond reload
# 查看 ntpdate 命令绝对路径 which ntpdate
