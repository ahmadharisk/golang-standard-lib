package main

import (
	"encoding/csv"
	"os"
)

func main() {
	csvWriter := csv.NewWriter(os.Stdout)
	_ = csvWriter.Write([]string{"jerapah", "kalkun", "lampu"})
	_ = csvWriter.Write([]string{"macan", "nasi", "orang"})
	_ = csvWriter.Write([]string{"pakan", "queen", "ratu"})
	csvWriter.Flush()
}
