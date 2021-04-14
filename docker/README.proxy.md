## 国内镜像代理
>- https://docs.docker.com/config/daemon/systemd/#httphttps-proxy

```
mkdir -p /etc/systemd/system/docker.service.d && cd /etc/systemd/system/docker.service.d

cat <<EOF >http-proxy.conf
[Service]
Environment="HTTPS_PROXY=https://registry.docker-cn.com:443"
Environment="HTTP_PROXY=http://hub-mirror.c.163.com:80"
Environment="NO_PROXY=localhost,127.0.0.1"
EOF 

```

>- systemctl daemon-reload
>- systemctl restart docker
>- docker info  可见 HTTP Proxy | HTTPS Proxy | No Proxy
>- 验证是否生效 systemctl show --property=Environment docker

## 国内镜像
>- Docker官方中国区 https://registry.docker-cn.com
>- 网易 http://hub-mirror.c.163.com
>- 中国科学技术大学 https://docker.mirrors.ustc.edu.cn
>- 阿里云Docker镜像加速器 需要登录 https://cr.console.aliyun.com/undefined/instances/mirrors
