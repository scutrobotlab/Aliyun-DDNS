package internal

import (
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type rr_domain_type struct {
	RR     string
	Domain string
	Type   string
}
type id_addr struct {
	ID   string
	Addr string
}

// Update DNS record with aliyun API. / 使用阿里云 API 更新 DNS 记录。
func UpdateRecord(client *alidns.Client, config []Config, ipv4_addr map[string]string, ipv6_addr map[string]string) {
	domains := map[string]bool{}
	for _, c := range config {
		domains[c.Domain] = true
	}

	records := map[rr_domain_type]id_addr{}

	for d := range domains {
		request := alidns.CreateDescribeDomainRecordsRequest()
		request.Scheme = "https"
		request.DomainName = d
		response, err := client.DescribeDomainRecords(request)
		if err != nil {
			log.Panicln(err.Error())
		}
		for _, r := range response.DomainRecords.Record {
			records[rr_domain_type{
				RR:     r.RR,
				Domain: d,
				Type:   r.Type,
			}] = id_addr{
				ID:   r.RecordId,
				Addr: r.Value,
			}
		}
	}

	for _, conf := range config {
		var (
			addr string
			ok   bool
		)
		log.Println("Checking record. / 检查记录。", conf)
		switch conf.Type {
		case "AAAA":
			addr, ok = ipv6_addr[conf.Interface]
		case "A":
			addr, ok = ipv4_addr[conf.Interface]
		default:
			log.Println("Invalid record type, skipped. / 不合法的记录类型。", conf.Type)
			continue
		}
		if !ok {
			log.Println("Interface IP unfound, skipped. / 找不到接口对应IP。", conf.Interface)
			continue
		}
		rdt := rr_domain_type{
			RR:     conf.RR,
			Domain: conf.Domain,
			Type:   conf.Type,
		}
		rec, ok := records[rdt]
		if !ok {
			log.Println("Domain record not found. / 未找到匹配的记录。", rdt)
			continue
		}
		if rec.Addr == addr {
			log.Println("Domain record don't need to be modified. / 域名记录无需修改。IP: ", addr)
			continue
		}
		request := alidns.CreateUpdateDomainRecordRequest()
		request.RecordId = rec.ID
		request.RR = conf.RR
		request.Type = conf.Type
		request.Value = addr
		response, err := client.UpdateDomainRecord(request)
		if err != nil {
			log.Panicln(err.Error())
		}
		log.Println("Domain record has been modified. / 域名记录已修改。IP: ", addr, response)

	}
}
