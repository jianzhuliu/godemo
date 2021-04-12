#!/bin/bash

echo building server:build
## 创建一个临时镜像
docker build -t server:build . -f Dockerfile.build

## 构建一个临时容器
docker create --name builder server:build

##从上一个临时容器中复制二进制文件到当前文件夹下
docker cp builder:/go/src/gitee.com/jianzhuliu/docker/server ./server

## 删除临时容器
docker rm -f builder

## 删除临时镜像
docker image rm server:build

echo building server:latest

##构建最终镜像
docker build -t server:latest .

## 删除本地二进制文件

## 测试运行镜像 
##docker run --rm --name=server -p 8083:8003 server:latest


rm server
