# Aliyun DDNS / 阿里云动态DNS

[![Go](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/release.yml/badge.svg)](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/release.yml)
[![Docker](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/docker.yml/badge.svg)](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/docker.yml)

Light-weight dynamic DNS tool for aliyun based on Go.

基于Go的轻量阿里云动态域名工具。

## Quick Start / 快速上手

二进制文件运行：

```bash
./aliyun-ddns 
    -r="cn-hangzhou"
    -d="www" 
    -n="example.com" 
    -i="em1"
    -a="/etc/AccessKey.csv"
```

## Flags Help / 参数帮助

- r  
    Region ID / 阿里云地域ID  
    参考: [阿里云地域ID文档](https://help.aliyun.com/document_detail/40654.html)  
    默认值：`cn-hangzhou`
- n  
    Domain Name / 域名  
	如 www.example.com , 此处填写 `example.com`  
    默认值：无
- d  
    Domain RR keyword / 域名主机记录  
    如 www.example.com , 此处填写 `www`  
    默认值：无
- i  
    Network interface to bind / 绑定的网络接口  
    使用命令 `ip a` 查看  
    默认值：无
- a  
    File path of `AccessKey.csv` / `AccessKey.csv`文件路径  
    Reference / 参考: [阿里云AccessKey文档](https://help.aliyun.com/document_detail/38738.html)  
    默认值：`/etc/AccessKey.csv`

## About / 关于

Most IPv4 network operators provide dynamic IP (DHCP), and the IP address of the device changes every time it is connected to the network. This tool can automatically read the device IP and update the domain name resolution record. Other devices can resolve to the device IP through DNS domain name.

大多数IPv4网络运营商提供动态IP（DHCP），设备的IP地址在每次连接到网络时都会改变。这个工具可以自动读取设备IP并更新域名解析记录，其他设备可以通过DNS域名解析到该设备的IP。

You can use `crontab` to run this tool periodically for dynamic domain name record updates.

可使用 `crontab` 来定期运行此工具，实现动态域名记录更新。

```bash
contab -e
```

Write the following and press `z` twice to save. This configuration sets the task to execute every hour.

写入以下内容，并按两次`z`保存。该配置设定任务每小时执行。

```
0 * * * * ./aliyun-ddns -i="em1" -d="www" -n="example.com"
```
