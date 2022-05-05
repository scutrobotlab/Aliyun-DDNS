package main

import (
	"flag"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/scutrobotlab/aliyun-ddns/internal"
)

func main() {
	var (
		region_id         string
		domain_name       string
		domain_RRkeyword  string
		network_interface string
		accesskey_file    string
	)
	flag.StringVar(&region_id, "r", "cn-hangzhou", "Region ID / 阿里云地域ID [请参考: https://help.aliyun.com/document_detail/40654.html]")
	flag.StringVar(&domain_name, "n", "", "Domain Name / 域名 [如 www.example.com , 此处填写 example.com]")
	flag.StringVar(&domain_RRkeyword, "d", "", "Domain RR keyword / 域名主机记录 [如 www.example.com , 此处填写 www]")
	flag.StringVar(&network_interface, "i", "", "Network interface to bind / 绑定的网络接口 [使用命令 ip a 查看]")
	flag.StringVar(&accesskey_file, "a", "/etc/AccessKey.csv", "File path of `AccessKey.csv` / `AccessKey.csv`文件路径 [获取`AccessKey.csv`, 请参考: https://help.aliyun.com/document_detail/38738.html]")

	flag.Parse()

	accesskey := internal.ReadCSV(accesskey_file)

	client, err := alidns.NewClientWithAccessKey(region_id, accesskey.ID, accesskey.Secret)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Bound network interface / 绑定的网络接口为: ", network_interface)
	addr := internal.GetInterfaceIPv4Addr(network_interface)
	log.Println("Got IPv4 address / 查询到IPv4地址: ", addr)

	internal.UpdateRecord(client, domain_RRkeyword, domain_name, addr)
}
