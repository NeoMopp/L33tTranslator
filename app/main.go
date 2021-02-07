package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Operation string

const (
    Leet Operation = "l33t"
    Md5 Operation = "md5"
    Md4 Operation = "md4"
    Sha1 Operation = "sha1"
    Sha256 Operation = "sha256"
    Sha384 Operation = "sha384"
    Sha512 Operation = "sha512"
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

func parseArguments() (word string, operation Operation) {
	wordPtr := flag.String("word", "foo", "the string you want converted")
	operationPtr := flag.String("operation", "default", "the operation you want to perform on the string")
	flag.Parse()
	//Can't ever be null as far as I can tell
	return *wordPtr, Operation(*operationPtr)
}

func main() {
	word, operation := parseArguments()

	fmt.Println(operation)

    //Is there a better way of doing this? Returning the enum before?
    switch operation {
    case Leet:
        fmt.Println(Leet)
    case Md5:
        fmt.Println(Md5)
    case Md4:
    case Sha1:
    case Sha256:
    case Sha384:
    case Sha512:
    default:
        fmt.Println("Not a valid Operation: " + operation)
    }


	recs := readCsv("matrix/l33t.csv")

	translator := csvRecordsToMap(recs)
	result := ""
	for _, c := range word {
		if len(translator[string(c)]) > 0 {
			result = result + translator[string(c)]
		} else {
			result = result + string(c)
		}
	}

	fmt.Println(result)
}
