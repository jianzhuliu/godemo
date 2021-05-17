## ClickHouse 
>- https://clickhouse.tech/

## 检查是否支持安装
>- grep -q sse4_2 /proc/cpuinfo && echo "SSE 4.2 supported" || echo "SSE 4.2 not supported"

## install
```
yum install -y yum-utils
rpm --import https://repo.clickhouse.tech/CLICKHOUSE-KEY.GPG
yum-config-manager --add-repo https://repo.clickhouse.tech/rpm/clickhouse.repo
yum install -y clickhouse-server clickhouse-client

/etc/init.d/clickhouse-server start
clickhouse-client

```

## docker
```
docker pull yandex/clickhouse-server:21.3.11.5
docker run -d --name some-clickhouse-server --ulimit nofile=262144:262144 yandex/clickhouse-server:21.3.11.5


docker pull yandex/clickhouse-client:21.3.11.5
docker run -it --rm --link some-clickhouse-server:clickhouse-server yandex/clickhouse-client:21.3.11.5 --host clickhouse-server

```

## 以上都失败了，考虑源码安装
>- https://clickhouse.tech/docs/zh/getting-started/install/#from-sources
>- https://clickhouse.tech/docs/zh/development/build/
