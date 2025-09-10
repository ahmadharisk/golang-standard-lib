package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "ayam, bebek, cicak\n" +
		"dadu, efek, fikri\n" +
		"gajah, harimau, ikan"

	csvReader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(record)
	}
}
