# blog-api


## 编译程序应用


简单编译

```shell
go build
```

编译linux系统下的二进制文件
```shell
CG0_ENABLED=0 GOODS=linux go build -a -o blog-api .
```

## 部署应用

nginx.conf添加代理，例：

```shell
server {
       listen       80;
       server_name  blog-api.test;

       location / {
           proxy_pass http://127.0.0.1:8000/;
       }
   }
```