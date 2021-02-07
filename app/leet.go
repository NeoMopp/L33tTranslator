package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func csvRecordsToMap(records [][]string) map[string]string {
	m := make(map[string]string)
	for i := 0; i < len(records); i += 1 {
		m[records[i][0]] = records[i][1]
	}
	return m
}

func leetHash(input string) string {
	recs := readCsv("matrix/l33t.csv")

	translator := csvRecordsToMap(recs)
	var result string
	for _, c := range input {
		if len(translator[string(c)]) > 0 {
			result = result + translator[string(c)]
		} else {
			result = result + string(c)
		}
	}
	return result
}
