## grpc
>- https://grpc.io/docs/languages/go/quickstart/
>- https://github.com/grpc/grpc-go
>- https://github.com/golang/protobuf
>- https://github.com/protocolbuffers/protobuf/releases

## go 环境配置
>- 支持 gomodule 及 国内代理
```
export GO111MODULE=on
export GOPROXY='https://goproxy.cn,direct'


```

## 下载安装 protoc  
```
yum install -y unzip
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d protoc-3.15.8
install protoc-3.15.8/bin/protoc /usr/local/bin/protoc
protoc --version

```

## 安装 protoc-gen-go | protoc-gen-go-grpc
>- go get -v google.golang.org/protobuf/cmd/protoc-gen-go
>- go get -v google.golang.org/grpc/cmd/protoc-gen-go-grpc

## example
```
wget https://github.com/grpc/grpc-go/archive/refs/tags/v1.37.0.tar.gz
tar -zxf v1.37.0.tar.gz
ls -l grpc-go-1.37.0/examples/helloworld/

```

## helloworld
```
cd helloworld
protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	helloworld/helloworld.proto


``` 

>- go run greeter_server/main.go
>- go run greeter_client/main.go jianzhu


