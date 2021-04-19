## Deployment
>- https://kubernetes.io/zh/docs/tasks/run-application/run-stateless-application-deployment/

#### 准备 deployment yaml 文件
>- wget https://k8s.io/examples/application/deployment.yaml
>- wget https://k8s.io/examples/application/deployment-update.yaml
>- wget https://k8s.io/examples/application/deployment-scale.yaml

#### 查看状态
>- minikube status    
>- kubectl get nodes
>- kubectl get deployments

#### 开启监听 deployment
>- kubectl get --watch deployment

#### 运行一个应用
>- kubectl apply -f deployment.yaml
>- kubectl describe deployment nginx-deployment
>- kubectl get pods -l app=nginx

#### 水平扩展
>- kubectl apply -f deployment-scale.yaml

#### 版本更新
>- kubectl apply -f deployment-update.yaml

### 删除 Deployment
>- kubectl delete deployment nginx-deployment