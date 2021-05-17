#!/bin/bash

yum install -y yum-utils
rpm --import https://repo.clickhouse.tech/CLICKHOUSE-KEY.GPG
yum-config-manager --add-repo https://repo.clickhouse.tech/rpm/clickhouse.repo
yum install -y clickhouse-server clickhouse-client

/etc/init.d/clickhouse-server start
clickhouse-client
