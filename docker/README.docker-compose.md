## docker-compose
>- https://docs.docker.com/compose/install/

## 安装
>- curl -L "https://github.com/docker/compose/releases/download/1.29.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
>- chmod +x /usr/local/bin/docker-compose
>- docker-compose --version

## 使用
>- https://github.com/docker/awesome-compose

## 操作命令
>- docker-compose [-f <arg>...] [options] [--] [COMMAND] [ARGS...]

>- options
```
  -f, --file FILE             指定 docker-compose 文件，默认为 docker-compose.yml
  -p, --project-name NAME     指定项目名称，默认使用当前目录名称作为项目名称
  --verbose                   输出调试信息
  --log-level LEVEL           日志级别 (DEBUG, INFO, WARNING, ERROR, CRITICAL)
  -v, --version               输出当前版本并退出
  -H, --host HOST             指定要连接的 Docker 地址
  --tls                       启用 TLS 认证
  --tlscacert CA_PATH         TLS CA 证书路径
  --tlscert CLIENT_CERT_PATH  TLS 公钥证书问价
  --tlskey TLS_KEY_PATH       TLS 私钥证书文件
  --tlsverify                 使用 TLS 校验对端
  --skip-hostname-check       不校验主机名
  --project-directory PATH    指定工作目录，默认是 Compose 文件所在路径。
```

>- COMMAND

```
  build              构建服务
  config             校验和查看 Compose 文件
  create             创建服务
  down               停止服务，并且删除相关资源
  events             实时监控容器的时间信息
  exec               在一个运行的容器中运行指定命令
  help               获取帮助
  images             列出镜像
  kill               杀死容器
  logs               查看容器输出
  pause              暂停容器
  port               打印容器端口所映射出的公共端口
  ps                 列出项目中的容器列表
  pull               拉取服务中的所有镜像
  push               推送服务中的所有镜像
  restart            重启服务
  rm                 删除项目中已经停止的容器
  run                在指定服务上运行一个命令
  scale              设置服务运行的容器个数
  start              启动服务
  stop               停止服务
  top                限制服务中正在运行中的进程信息
  unpause            恢复暂停的容器
  up                 创建并且启动服务
  version            打印版本信息并退出
```