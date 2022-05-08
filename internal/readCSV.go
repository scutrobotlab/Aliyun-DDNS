package internal

import (
	"encoding/csv"
	"log"
	"os"
)

type AccessKey struct {
	ID     string
	Secret string
}

// Read CSV file and get Access Key. / 从 csv 文件中获取 Access Key。
func GetAccessKey(file string) AccessKey {
	f, err := os.Open(file)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Panicln(err.Error())
	}

	return AccessKey{
		ID:     data[1][0],
		Secret: data[1][1],
	}
}

type Config struct {
	RR        string
	Domain    string
	Type      string
	Line      string
	Interface string
}

// Read CSV file and get Config. / 从 csv 文件中获取配置。
func GetConfig(file string) []Config {
	f, err := os.Open(file)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Panicln(err.Error())
	}

	header := map[string]int{}
	for i, s := range data[0] {
		header[s] = i
	}

	config := make([]Config, len(data)-1)
	for i := range data {
		if i == 0 {
			continue
		}
		config[i].RR = data[i][header["RR"]]
		config[i].Domain = data[i][header["Domain"]]
		config[i].Type = data[i][header["Type"]]
		config[i].Line = data[i][header["Line"]]
		config[i].Interface = data[i][header["Interface"]]
	}

	return config
}
