## Kubernetes 
>- https://kubernetes.io/
>- https://kubernetes.io/zh/docs/home/

## minikute
>- https://minikube.sigs.k8s.io/docs/start/
>- 2 CPU及以上
>- 2GB 内存
>- 20GB 磁盘空间
>- 可联网
>- 安装有容器或者虚拟机，如Docker，VirtualBox, VMWare

## 下载
>- curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
>- sudo install minikube-linux-amd64 /usr/local/bin/minikube
>- or
>- curl -LO https://github.com/kubernetes/minikube/releases/download/v1.19.0/minikube-linux-amd64
>- install minikube-linux-amd64 /usr/local/bin
>- minikube version

## 配置
>- useradd kubedocker
>- passwd kubedocker
>- usermod -aG docker kubedocker
>- minikube config set driver docker
>- ln -s $(which minikube) /usr/local/bin/kubectl
>- kubectl version 

## 启动
>- su kubedocker
>- minikube start --help | grep mirror
>- minikube start --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'

>- 下次启动 
>- docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kicbase:v.0.0.20 kicbase/stable:v0.0.20
>- minikube start --driver=docker --base-image=kicbase/stable:v0.0.20

## 查看效果
>- minikube dashboard
>- or  minikube dashboard --url
>- 如果想虚拟机外访问，需要配置代理

```
nohup kubectl proxy  --port=8000 --address='192.168.126.71' --accept-hosts='^localhost$,^127\.0\.0\.1$,^192.168.126.71$'  >/dev/null 2>&1& 

```
>- 浏览器访问 192.168.126.71:8000
>- http://192.168.126.71:8000/version
>- http://192.168.126.71:8000/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/


## 部署应用

#### Deployment
>- kubectl create deployment hello-minikube --image=registry.cn-hangzhou.aliyuncs.com/google_containers/echoserver:1.4
>- kubectl get deployments
>- kubectl get pods
>- kubectl get events
>- kubectl config view 

#### Service
>- kubectl expose deployment hello-minikube --type=NodePort --port=8080
>- kubectl get services
>- minikube service hello-minikube

#### 端口转发
>- kubectl port-forward --address 0.0.0.0 service/hello-minikube 8002:8080
>- http://192.168.126.71:8002

## 清理
>- kubectl delete deployment hello-minikube
>- kubectl delete service hello-minikube

## 虚拟机处理
>- minikube stop
>- minikube delete