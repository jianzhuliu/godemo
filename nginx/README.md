## nginx
>- http://nginx.org/

## 安装
>- [install](./README.install.md)

## nginx.conf
>- [nginx.conf](./README.nginx.conf.md)

## 概念
>- 正向代理 ==> 为客户端服务，客户端通过代理访问无法正常访问到的服务器资源

>- 反向代理 ==> 为服务器服务，代理帮助服务器接收客户端连接
* 隐藏真实服务器
* 负载均衡便于横向扩充后端动态服务
* 动静分离，提升系统健壮性

>- 动静分离 ==> 一种将静态资源与动态资源分开不同系统访问的架构设计方式
* 静态资源，直接到静态资源目录获取
* 动态资源，通过反向代理，把请求转发给后台应用去处理

>- 负载均衡 ==> 利用服务器集群，将请求分发到不同的服务器，分担压力
* 轮询策略
* 最小连接数策略
* 最快响应时间策略
* 客户端 ip 绑定策略， 可以解决动态网页 session 共享问题
* 随机策略
* 哈希算法

## upstream
>- 定义上游服务器(后台提供的应用服务器)

```
upstream backend_server{
	server 192.168.126.71:8081 weight=3 max_conns=1000 fail_timeout=10s max_fails=2;
	keepalive 32;
	keepalive_requests 50;
	keepalive_timeout 30s;
}
```

#### 在 upstream 内可使用的命令
* server address [parameters]  ##定义上游服务器地址 
* zone  ##定义共享内存，用于跨 worker 子进程
* keepalive 16; ##每个 worker 子进程与上游服务器空闲长连接的最大数量
* keepalive_requests 100; ##一个长连接最多请求数
* keepalive_timeout 60s; ##空闲长连接的最长保持时间
* hash 哈希负载均衡算法
* ip_hash 依据 ip 进行哈希计算的负载均衡算法
* least_conn 最少连接数负载均衡算法
* least_time 最短响应时间负载均衡算法
* random 随机负载均衡算法

#### server 指令可选参数
>- weight=number 权重值，默认为1 
>- max_conns=number 上游服务器的最大并发连接数
>- fail_timeout=number 服务器不可用的判定时间
>- max_fails=number 服务器不可用的检查次数
>- backup 备份服务器，仅当其他服务器不可用时才会启用
>- down 标记服务器长期不可用，离线维护


## proxy_pass 
>- 用于配置代理服务器

```
语法：proxy_pass URL;

proxy_pass http://127.0.0.1:8081
```
 
#### URL 参数规则
>- 必须以 http 或 https 开头
>- 可以携带变量
>- 是否带 URI，会直接影响发往上游请求 URL

#### URL 不带 /
```
location /bbs/{
  proxy_pass http://127.0.0.1:8081;
}

``` 
>- 不带 / 意味着 Nginx 不会修改用户 URL ，而是直接透传给上游的应用服务器
>- 用户请求 URL ： /bbs/abc/test.html
>- 请求到达 Nginx 的 URL ： /bbs/abc/test.html 
>- 请求到达上游应用服务器的 URL ： /bbs/abc/test.html 

#### URL 带 /
```
location /bbs/{
  proxy_pass http://127.0.0.1:8081/;
}
```

>- 带 / 意味着 Nginx 会修改用户 URL ，修改方法是将 location 后的 URL 从用户 URL 中删除
>- 用户请求 URL ： /bbs/abc/test.html
>- 请求到达 Nginx 的 URL ： /bbs/abc/test.html 
>- 请求到达上游应用服务器的 URL ： /abc/test.html

## 配置反向代理
>- [proxy](./README.proxy.md)

## 配置负载均衡
>- [balance](./README.balance.md)

## 配置代理缓存
>- [cache](./README.cache.md)
