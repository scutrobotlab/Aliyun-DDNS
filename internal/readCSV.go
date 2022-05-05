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
func ReadCSV(file string) AccessKey {
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
