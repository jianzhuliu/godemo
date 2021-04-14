## etcd

## 安装 
>- [./install/etcd_bin_install.sh](./install/etcd_bin_install.sh)

## 集群搭建
>- 关闭及停用防火墙 systemctl stop firewalld , systemctl disable firewalld
>- 或者添加对应端口 firewall-cmd --add-port=2379-2380/tcp --permanent --zone=public
>- 配置3台机器ip及端口  [./install/etcd.common.sh](./install/etcd.common.sh)
>- 每台机器执行 [./install/etcd.run.sh](./install/etcd.run.sh)

## demo测试
>- 下载对应版本的 src,复制到对应gopath目录下
>- cd cmd && go run comm.go get.go