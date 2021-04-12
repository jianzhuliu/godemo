## Kubernetes
https://github.com/kubernetes/minikube

#### 配置条件
>- 2 CPU及以上
>- 2GB 内存
>- 20GB 磁盘空间
>- 可联网
>- 安装有容器或者虚拟机，如Docker，VirtualBox, VMWare

#### 安装 minikube
>- curl -LO https://github.com/kubernetes/minikube/releases/download/v1.19.0/minikube-linux-amd64
>- install minikube-linux-amd64 /usr/local/bin
>- https://minikube.sigs.k8s.io/docs/start/

#### 安装 kubectl
>- https://github.com/kubernetes/kubernetes/tree/master/CHANGELOG 查找对应的版本
>- curl -LO https://dl.k8s.io/v1.21.0/kubernetes-client-linux-amd64.tar.gz 
>- tar -xfv kubernetes-client-linux-amd64.tar.gz 
>- install kubernetes/client/bin/kubectl /usr/local/bin/kubectl

#### 安装集群
>- 非 root 用户运行
>- 创建用户 		useradd kubedocker
>- 设置密码 		passwd kubedocker   --123456
>- 添加到docker组 	usermod -aG docker kubedocker
>- 启动集群 		minikube start --driver=docker
>- 设置默认 		minikube config set driver docker
>- 如没有安装 kubectl,可设置

```
 alias kubectl="minikube kubectl --"
 ln -s $(which minikube) /usr/local/bin/kubectl
```

#### 检查集群状态
>- kubectl cluster-info

#### 创建一个应用
>- 创建部署文件 deployment.yaml
```
cat <<EOF >deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hello-world
  template:
    metadata:
      labels:
        app: hello-world
    spec:
      containers:
      - name: hello-world
        image: wilhelmguo/nginx-hello:v1
        ports:
        - containerPort: 80
EOF
```

>- 发布部署文件到集群
```
kubectl create -f deployment.yaml
```

>- 检查 Pod 是否启动成功
```
kubectl get pod -o wide
```

>- 创建服务文件 service.yaml
```
cat <<EOF >service.yaml
apiVersion: v1
kind: Service
metadata:
  name: hello-world
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 80
  selector:
    app: hello-world
EOF
```

>- 创建 Service
```
kubectl create -f service.yaml
```
	
>- 查看服务信息
```
kubectl get service -o wide
```	
		
>- 对外访问
```
minikube service hello-world
```

>- curl http://192.168.49.2:30928

