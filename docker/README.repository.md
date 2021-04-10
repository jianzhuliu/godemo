# 仓库(Repository)
>- 仓库是分发和存储 docker 镜像的地方

## 公共镜像仓库
>- 创建账号 https://hub.docker.com/
>- 拉取镜像 docker pull busybox
>- 登陆公共镜像仓库 docker login
>- 镜像重命名 docker tag busybox jianzhuliu/busybox
>- 上传镜像 docker push jianzhuliu/busybox

## 搭建私有仓库
>- 启动本地仓库 docker run -d -p 5000:5000 --name registry registry
>- 镜像重命名 docker tag busybox localhost:5000/busybox
>- 上传镜像 docker push localhost:5000/busybox
>- 删除相关镜像 docker rmi busybox localhost:5000/busybox
>- 查看是否成功删除 docker image ls
>- 本地仓库拉取镜像 docker pull localhost:5000/busybox

>- 持久化镜像存储  docker run -v /data/docker/data:/var/lib/registry -d -p 5000:5000 --name registry registry
>- 参数配置文档  https://github.com/distribution/distribution/blob/main/docs/configuration.md
>- 删除镜像  docker stop registry && docker rm registry

## 构建外部可访问的仓库
>- 需要合法的域名及CA证书 

```
$ docker run -d \
  --name registry \
  -v /local/path/data:/var/lib/registry \
  -v /local/path/certs:/certs \
  -e REGISTRY_HTTP_ADDR=0.0.0.0:443 \
  -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/regisry.crt \
  -e REGISTRY_HTTP_TLS_KEY=/certs/regisry.key \
  -p 443:443 \
  registry:2.7


```