# Aliyun DDNS / 阿里云动态DNS

Light-weight dynamic DNS tool for aliyun based on Go.  
基于Go的轻量阿里云动态域名工具。

## Quick Start / 快速上手

```bash
./aliyun-ddns 
    -r="cn-hangzhou"
    -d="www" 
    -n="example.com" 
    -i="em1"
    -a="/etc/AccessKey.csv"
```

## About Flags / 关于参数

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