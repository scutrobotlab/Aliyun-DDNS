# Aliyun DDNS / 阿里云动态DNS

[![Go](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/release.yml/badge.svg)](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/release.yml)
[![Docker](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/docker.yml/badge.svg)](https://github.com/scutrobotlab/Aliyun-DDNS/actions/workflows/docker.yml)

基于Go的轻量阿里云动态域名工具。

Light-weight dynamic DNS tool for aliyun based on Go.

## 快速上手

1. 从阿里云下载`AccessKey.csv`。  

   > 参考 [阿里云AccessKey文档](https://help.aliyun.com/document_detail/38738.html) 

2. 配置`Config.csv`。

    ```csv
    RR,Domain,Type,Line,Interface
    *,example.com,A,default,em1
    @,example.com,A,default,em1
    www,example.com,A,default,em1
    ```

    CSV文件可以使用Excel等表格工具查看。如示例的csv以表格形式展示为：

    | RR  | Domain      | Type | Line    | Interface |
    | --- | ----------- | ---- | ------- | --------- |
    |  *  | example.com | A    | default | em1       | 
    |  @  | example.com | A    | default | em1       | 
    | www | example.com | A    | default | em1       | 

    > 其中 `Interface` 为绑定的网卡名称，可以用命令`ip a`查看。  
    > 程序将读取网卡的当前IP，以更新对应的域名解析记录。  
    > 程序将根据 `Type` 自动区分 IPv4 和 IPv6。  
    > 参考 [阿里云控制台 云解析DNS/域名解析/解析设置](https://dns.console.aliyun.com/#/dns/domainList)

3. 二进制文件运行：

    ```bash
    ./aliyun-ddns 
        -r="cn-hangzhou"
        -a="./AccessKey.csv"
        -c="./AccessKey.csv"
    ```

## 运行参数

| 参数名  |  注释                                   | 默认值            |
| ------ | --------------------------------------- | ---------------- |
|  r     | [阿里云地域ID](https://help.aliyun.com/document_detail/40654.html) | `cn-hangzhou` |
|  a     | `AccessKey.csv`文件路径                       | `/etc/AccessKey.csv`       | 
|  c     | `Config.csv`文件路径                          | `/etc/Config.csv`       | 

## About / 关于

大多数IPv4网络运营商提供动态IP（DHCP），设备的IP地址在每次连接到网络时都会改变。这个工具可以自动读取设备IP并更新域名解析记录，其他设备可以通过DNS域名解析到该设备的IP。

Most IPv4 network operators provide dynamic IP (DHCP), and the IP address of the device changes every time it is connected to the network. This tool can automatically read the device IP and update the domain name resolution record. Other devices can resolve to the device IP through DNS domain name.

可使用 `crontab` 来定期运行此工具，实现动态域名记录更新。

You can use `crontab` to run this tool periodically for dynamic domain name record updates.

```bash
contab -e
```

写入以下内容，并按两次`z`保存。该配置设定任务每小时执行。

Write the following and press `z` twice to save. This configuration sets the task to execute every hour.

```
0 * * * * aliyun-ddns
```
