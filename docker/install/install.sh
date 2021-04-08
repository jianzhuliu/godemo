#/bin/bash

## 更新 yum 
yum update

## 配置镜像源
wget -O /etc/yum.repos.d/docker-ce.repo https://download.docker.com/linux/centos/docker-ce.repo
sed -i 's+download.docker.com+mirrors.cloud.tencent.com/docker-ce+' /etc/yum.repos.d/docker-ce.repo
yum makecache fast

## 查看仓库中所有docker版本
yum list docker-ce --showduplicates | sort -r

## 安装
yum -y install docker-ce

## 指定版本安装
##yum -y install docker-ce-3:20.10.5-3.el7

## 查看版本
docker version

#### systemctl
## 启动
systemctl start docker

## 重启
systemctl restart docker

## 加入开机启动
systemctl enable docker

## 测试 
docker run hello-world

## 卸载 
#yum remove docker-ce 

## 删除所有镜像，容器和存储卷
#rm -rf /var/lib/docker 



#######镜像安装
#wget -O /tmp https://download.docker.com/linux/centos/7/x86_64/stable/Packages/docker-ce-20.10.5-3.el7.x86_64.rpm
#wget -O /tmp https://mirrors.cloud.tencent.com/docker-ce/linux/centos/7/x86_64/stable/Packages/docker-ce-20.10.5-3.el7.x86_64.rpm

#yum install /tmp/docker-ce-20.10.5-3.el7.x86_64.rpm


####运行官网 tutorial 
docker -d -p 2345:80 docker/getting-started

## 参数说明
#-d	后台运行
#-p 2345:80 端口映射，主机端口(2345):容器内端口(80)
#docker/getting-started	镜像


## 浏览器访问
#http://192.168.195.71:2345/

