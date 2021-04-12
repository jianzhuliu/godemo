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
