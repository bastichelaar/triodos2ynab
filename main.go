package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var file string
	var outfile string

	flag.StringVar(&file, "in", "mutations.csv", "File to convert")
	flag.StringVar(&outfile, "out", "ynab.csv", "Output file")
	flag.Parse()

	csvfile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	targetcsv, err := os.Create(outfile)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer targetcsv.Close()

	writer := csv.NewWriter(targetcsv)
	header := []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"}
	writer.Write(header)

	for _, row := range rawCSVdata {
		date := strings.Replace(row[0], "-", "/", -1)
		payee := row[4]
		category := ""
		memo := row[7]
		inflow := ""
		outflow := ""

		amount := strings.Replace(row[2], ",", ".", -1)
		if row[3] == "Credit" {
			inflow = amount
		} else {
			outflow = amount
		}

		writer.Write([]string{date, payee, category, memo, outflow, inflow})
		fmt.Printf("Date: %s --- Name: %s\n", row[0], row[1])
	}
}
