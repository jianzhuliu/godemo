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
```
server {
	listen [::]:80;
	listen [::]:443 ssl http2;   # SSL 访问端口号为 443
	server_name www.example.com;         # 填写绑定证书的域名(我这里是随便写的)
	ssl_certificate /etc/nginx/https/www.example.com.crt;   # 证书地址
	ssl_certificate_key /etc/nginx/https/www.example.com.key;      # 私钥地址
	#ssl_session_timeout 10m;
	#ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # 支持ssl协议版本，默认为后三个，主流版本是[TLSv1.2]

	location / {
	root         /usr/share/nginx/html;
	index        index.html index.htm;
	}
}


```