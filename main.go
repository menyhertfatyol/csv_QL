package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filePaths := []string{}
	patterns := []Pattern{}

	for _, arg := range os.Args[1:] {
		if isExists(arg) {
			filePaths = append(filePaths, arg)
		} else {
			pattern, err := newPattern(arg)
			if err != nil {
				log.Fatal("The following pattern is invalid " + arg + "\n" + err.Error())
			}
			patterns = append(patterns, pattern)
		}
	}

	for _, path := range filePaths {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		if err := filterCSVFile(file, patterns); err != nil {
			log.Fatal(err)
		}
	}
}

func filterCSVFile(file *os.File, patterns []Pattern) error {
	readMyCsv := csv.NewReader(file)
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		record, err := readMyCsv.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		for _, pattern := range patterns {
			if pattern.Expression.MatchString(record[pattern.ColumnNumber-1]) {
				writer.Write(record)
				break
			}
		}
	}

	return nil
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Pattern ...
type Pattern struct {
	ColumnNumber int
	Expression   *regexp.Regexp
}

func newPattern(s string) (Pattern, error) {
	format := regexp.MustCompile("[1-9]+:.*")
	if !format.MatchString(s) {
		return Pattern{}, errors.New("Invalid pattern format")
	}
	parts := strings.SplitN(s, ":", 2)
	columnNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return Pattern{}, errors.New("Column number must be an integer: " + parts[0])
	}
	expression, err := regexp.Compile(regexp.QuoteMeta(parts[1]))
	if err != nil {
		return Pattern{}, err
	}
	return Pattern{ColumnNumber: columnNumber, Expression: expression}, nil
}
