## 镜像构建
>- Docker 镜像是分层的，并且每一层镜像都会额外占用存储空间，一个镜像层数越多，占用存储空间就越多
>- 构建镜像要尽量保持镜像体积小
>- 1、基础镜像体积尽量小
>- 2、尽量减少 Dockerfile 的命令行数, 每一条指令都会生成一个镜像层

## go demo | main.go
>- 编译构建二进制对象
>- 运行构建镜像

```
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}

func main() {
	addr := ":8003"
	fmt.Printf("going to serve %s\n", addr)
	http.HandleFunc("/", hello)
	if err:= http.ListenAndServe(addr, nil); err != nil {
	fmt.Println(err)
	}
}

```

## 非多阶段构建
>- Dockerfile.build
```
FROM golang:1.14.15-alpine3.12
WORKDIR /go/src/gitee.com/jianzhuliu/docker
COPY main.go . 
RUN go build -o server .
```

>- Dockerfile
```
FROM alpine:3.12
WORKDIR /root/
COPY server .
CMD ["./server"] 

```

>- 辅助脚本 build.sh 
```
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
rm server

## 测试运行镜像 
##docker run --rm --name=server -p 8083:8003 server:latest

```

>- 查看镜像  docker image ls server
>- 运行镜像  docker run --rm --name=server -p 80:8003 server:latest
>- curl 127.0.0.1
>- 删除镜像 docker image rm server:latest

## 多阶段构建
>- 一个 Dockerfile 即可，不需要额外辅助脚本，镜像内支持使用多个 FROM,以最后一个 FROM 作为镜像的生成
>- 构建阶段支持命名 AS，如果不命名，默认第一个阶段序号为 0
>- Dockerfile
```
##第一阶段 AS 命名构建阶段，用于第二个阶段拷贝数据
FROM golang:1.14.15-alpine3.12 AS builder

##工作目录
WORKDIR /go/src/gitee.com/jianzhuliu/docker 

##复制本地文件到镜像中
COPY main.go .    

##执行编译
RUN go build -o server . 

##第二阶段构建
FROM alpine:3.12   

##工作目录
WORKDIR /root/  

##把第阶段构建的镜像中对应目录下二进制文件拷贝到当前镜像中
COPY --from=builder /go/src/gitee.com/jianzhuliu/docker/server .  

##执行命令
CMD ["./server"]  

```

>- 构建镜像  docker build -t server:multi .
>- 查看镜像  docker image ls server
>- 运行镜像  docker run --rm --name=server -p 8080:8003 server:multi
>- curl 127.0.0.1:8080
>- 删除镜像 docker image rm server:multi


## 停止在某个阶段
>- docker build --target=builder -t server:build .
>- 查看镜像  docker image ls server
>- 删除镜像  docker image rm server:build

## 使用现有镜像文件
>- COPY --from=nginx:latest /etc/nginx/nginx.conf /etc/local/nginx.conf