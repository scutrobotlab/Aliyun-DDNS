package main

import (
	"flag"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/scutrobotlab/aliyun-ddns/internal"
)

func main() {
	var (
		region_id      string
		accesskey_file string
		config_file    string
	)
	flag.StringVar(&region_id, "r", "cn-hangzhou", "Region ID / 阿里云地域ID [请参考: https://help.aliyun.com/document_detail/40654.html]")
	flag.StringVar(&accesskey_file, "a", "/etc/AccessKey.csv", "File path of `AccessKey.csv` / `AccessKey.csv`文件路径 [获取`AccessKey.csv`, 请参考: https://help.aliyun.com/document_detail/38738.html]")
	flag.StringVar(&config_file, "a", "/etc/Config.csv", "File path of `Config.csv` / `Config.csv`文件路径")

	flag.Parse()

	accesskey := internal.GetAccessKey(accesskey_file)

	client, err := alidns.NewClientWithAccessKey(region_id, accesskey.ID, accesskey.Secret)
	if err != nil {
		log.Panicln(err.Error())
	}

	ipv4_addr, ipv6_addr := internal.GetInterfaceAddrs()

	config := internal.GetConfig(config_file)

	internal.UpdateRecord(client, config, ipv4_addr, ipv6_addr)
}
