## https

## https 工作流程
>- 客户端请求百度 https://www.baidu.com
>- 百度给客户端下发 CA 证书
>- 客户端验证 CA 证书的有效合法性
>- 客户端验证通过后，生成一串随机数，然后使用公钥(CA 证书中的)加密后发给百度
>- 百度收到加密后的字符串，然后使用私钥解密，得到真实随机数
>- 百度使用随机数加密需要下发的内容给客户端
>- 客户端收到内容后，用随机数解密，得到真实内容 

## 配置证书
>- http://nginx.org/en/docs/http/configuring_https_servers.html

```
server {
	listen 80;
	listen 443 ssl;   # SSL 访问端口号为 443
	server_name www.example.com;         # 填写绑定证书的域名(我这里是随便写的)
	ssl_certificate /etc/nginx/https/www.example.com.crt;   # 证书地址
	ssl_certificate_key /etc/nginx/https/www.example.com.key;      # 私钥地址
	##ssl_ciphers         HIGH:!aNULL:!MD5; #加密算法
	#ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # 支持ssl协议版本，默认为后三个，主流版本是[TLSv1.2]

	location / {
	root         /usr/share/nginx/html;
	index        index.html index.htm;
	}
}


```

## demo
>- 生成证书
```
mkdir -p /tmp/nginx/{ca,html}
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /tmp/nginx/ca/ssl.key -out /tmp/nginx/ca/ssl.crt
```

>- 准备配置文件
```
cat >/tmp/nginx/default.conf<<EOF
server {
 listen 80;
 listen 443 default ssl;
 ssl_certificate /data/nginx/ca/ssl.crt;
 ssl_certificate_key /data/nginx/ca/ssl.key;
 server_name localhost;
 location / {
  root /usr/share/nginx/html;
  index index.html;
 }
}
EOF

cat >/tmp/nginx/html/index.html<<EOF
<h1>Welcome to use https</h1>
EOF


```

## 检查配置
```
docker run --rm --name=nginx -p 80:80 -p 443:443 \
-v /tmp/nginx/default.conf:/etc/nginx/conf.d/default.conf \
-v /tmp/nginx/html:/usr/share/nginx/html \
-v /tmp/nginx/ca:/data/nginx/ca \
nginx:1.20.0 nginx -t
```

## 前台运行
```
docker run --rm --name=nginx -p 80:80 -p 443:443 \
-v /tmp/nginx/default.conf:/etc/nginx/conf.d/default.conf \
-v /tmp/nginx/html:/usr/share/nginx/html \
-v /tmp/nginx/ca:/data/nginx/ca \
nginx:1.20.0 nginx -g "daemon off;"
```

## 浏览器访问，比如
>- https://192.168.126.71/
>- http://192.168.126.71/