#/bin/bash

## win虚拟机clone之后初始化机器

help() {
	echo "Usage: ./clone_init.sh centos6_02 62"
	echo "param1: hostname"
	echo "param2: last ip of IPADDR"
	echo "ERROR: $1. Aborting!"
	exit 1
}

if [ "$1" == "" ];then
	help "param1 is not be empty"
fi

if [ "$2" == "" ];then
	help "param2 is not be empty"
fi

if ! echo $2 | egrep -q "^[0-9]+$"
then 
	help "param2 should be number"
fi

name=$1
ipnum=$2

echo "going to exec ./clone_init.sh $name $ipnum"

#修正主机名
#sed "s/\(HOSTNAME=\).*/\1centos/" /etc/sysconfig/network
sed -i "s/\(HOSTNAME=\).*/\1$name/" /etc/sysconfig/network

#修改ip 
#sed "s/\(IPADDR=\([0-9]\?[0-9][0-9]\?\.\)\{3\}\).*/\134/" /etc/sysconfig/network-scripts/ifcfg-eth0
sed -i "s/\(IPADDR=\([0-9]\?[0-9][0-9]\?\.\)\{3\}\).*/\1$ipnum/" /etc/sysconfig/network-scripts/ifcfg-eth0


#删除文件
rm -f /etc/udev/rules.d/70-persistent-net.rules

#重启网卡
#service network restart

#重启机器
reboot
