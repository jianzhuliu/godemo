#!/bin/bash

## 创建一个 127.0.0.1:6379 的 tcp 连接
exec 8<> /dev/tcp/127.0.0.1/6379

## 发送指令，redis 删除 key 操作到文件描述符中
echo "del geos" >& 8

for ((i=1;i<5; i++)) do
	for ((j=1; j<5;j++)) do
		## 创建点坐标
		echo "geoadd geos $i $j $i.$j" 
		echo "geoadd geos $i $j $i.$j" >& 8
	done
done

## 关闭文件描述符
exec 8<&-

