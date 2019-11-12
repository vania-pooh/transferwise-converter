package main

import (
	"encoding/csv"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		inputFile := os.Args[1]
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("failed to read input file %s: %v", inputFile, err)
		}
		r := csv.NewReader(strings.NewReader(string(data)))
		records, err := r.ReadAll()
		if err != nil {
			log.Fatalf("failed to parse input file %s: %v", inputFile, err)
		}
		var converted [][]string
		converted = append(converted, []string{"DATE", "BANK TRANSACTION ID", "BUSINESS PARTNER NAME", "PAYMENT DETAILS", "AMOUNT"})
		for _, row := range records {
			if len(row) >= 11 {
				amount := row[2]
				if n, err := strconv.ParseFloat(amount, 32); err != nil || n < 0 {
					continue
				}
				newRow := []string{
					strings.Replace(row[1], "-", ".", -1), row[0], row[10], row[4], amount,
				}
				converted = append(converted, newRow)
			}
		}

		w := csv.NewWriter(os.Stdout)
		for _, cr := range converted {
			if err := w.Write(cr); err != nil {
				log.Fatalf("error writing record to csv: %v\n", err)
			}
		}

		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("input file not specified")
	}
}
