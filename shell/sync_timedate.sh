#/bin/bash

#设置硬件时钟与本地时钟一致
timedatectl set-local-rtc 1

#设置时区
timedatectl set-timezone Asia/Shanghai

# centos7 /etc/localtime 记录本金时间信息 /usr/share/zoneinfo 目录下存放不同时区信息文件
# 构建软连接即可 ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 不能用cp来复制文件，否则会将localtime文件原来所链接的文件修改


#安装 ntpdate 
# Network Time Protocol
yum -y install ntpdate

# 同步时间
# pool.ntp.org
# cn.ntp.org.cn
# hk.ntp.org.cn
# us.ntp.org.cn

## 用网络服务器时间同步操作系统时间
ntpdate -u cn.ntp.org.cn

## 用操作系统时间同步硬件时间
hwclock -w


## 查看系统时间
date

## 查看硬件时间
hwclock -r 

# 一般使用 crontab 来同步时间
# yum -y install crontab
# crontab -e
# */5 * * * * /usr/sbin/ntpdate cn.ntp.org.cn
# service crond reload
# 查看 ntpdate 命令绝对路径 which ntpdate
