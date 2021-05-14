#!/bin/bash

## 创建一个 127.0.0.1:6379 的 tcp 连接
exec 10<> /dev/tcp/127.0.0.1/6379

for i in {20210513..20210514}; do
	## 删除指定的key
	echo "del $i" >& 10
	## 随机登录数
	k=$[RANDOM/10000+1]
	for ((j=1;j<=$k;j++)) do
		## 设置随机登录的角色id
		rid=$[RANDOM/10000+2]
		## 设置登录情况
		echo "setbit $i $rid 1" 
		echo "setbit $i $rid 1" >& 10
	done
done

## 关闭文件描述符
exec 10<&-

