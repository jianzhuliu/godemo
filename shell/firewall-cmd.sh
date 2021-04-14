#!/bin/bash

###防火墙  centos7.9

## 开机开启及停止防火墙 
systemctl enable firewalld
systemctl disable firewalld

## 开启及停止防火墙
systemctl start firewalld
systemctl stop firewalld

## 查看状态
systemctl status firewalld
systemctl is-enabled firewalld


## 查看所有打开的端口
firewall-cmd --list-ports

## 显示状态
firewall-cmd --state

## 添加访问端口
firewall-cmd --zone=public --add-port=8004/tcp --permanent
firewall-cmd --zone=public --add-port=8003/tcp --permanent
firewall-cmd --zone=public --add-port=8002/tcp --permanent
firewall-cmd --zone=public --add-port=8001/tcp --permanent
firewall-cmd --zone=public --add-port=8000/tcp --permanent

## 更新防火墙 
firewall-cmd --reload