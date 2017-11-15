package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	myRegex := "hello"
	regexToFind, err := regexp.Compile(myRegex)
	check(err)
	csvFile, err := os.Open("test.csv")
	check(err)
	readMyCsv := csv.NewReader(csvFile)
	for {
		record, err := readMyCsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if regexToFind.MatchString(record[0]) {
			fmt.Println(record)
		}

	}
}
