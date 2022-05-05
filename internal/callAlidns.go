package internal

import (
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// Update DNS record with aliyun API. / 使用阿里云 API 更新 DNS 记录。
// Check if domain record value same as addr. If not, modified it. / 检查 domain 记录值与 addr 是否一致。如不一致，则修改记录值。
func UpdateRecord(client *alidns.Client, domain_RRkeyword string, domain_name string, addr string) {
	var record_id string
	var record_rr string
	var record_type string
	var record_addr string

	log.Println("Checking domain / 检查域名: ", domain_RRkeyword+domain_name)

	{
		request := alidns.CreateDescribeDomainRecordsRequest()
		request.Scheme = "https"
		request.DomainName = domain_name
		request.RRKeyWord = domain_RRkeyword
		response, err := client.DescribeDomainRecords(request)
		if err != nil {
			log.Panicln(err.Error())
		}
		for _, r := range response.DomainRecords.Record {
			record_id = r.RecordId
			record_rr = r.RR
			record_type = r.Type
			record_addr = r.Value
		}
	}

	if record_addr != addr {
		request := alidns.CreateUpdateDomainRecordRequest()
		request.RecordId = record_id
		request.RR = record_rr
		request.Type = record_type
		request.Value = addr
		response, err := client.UpdateDomainRecord(request)
		if err != nil {
			log.Panicln(err.Error())
		}
		log.Println("Domain record has been modified. / 域名记录已修改。", response.String())
	} else {
		log.Println("Domain record don't need to be modified. / 域名记录无需修改。")
	}
}
