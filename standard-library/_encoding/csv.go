package _encoding

import (
	"encoding/csv"
	"os"
)

func ExportToCsv(filename string, records [][]string) error {
	csvf, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer csvf.Close()

	w := csv.NewWriter(csvf)
	return w.WriteAll(records)
}

func ReadFromCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	w := csv.NewReader(file)
	return w.ReadAll()
}
