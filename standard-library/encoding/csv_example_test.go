package encoding

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvFileWriter(t *testing.T) {
	records := [][]string{
		{"Name", "Age", "Description"},
		{"maodou", "27", "actor"},
		{"andy", "20", "to"},
	}

	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	assert.NoError(t, err)
}

func TestCsvFileReader(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	w := csv.NewReader(file)
	records, err := w.ReadAll()
	if err != nil {
		t.Error(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}

func TestCsvParseFile(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	type User struct {
		Name        string
		Age         string
		Description string
	}

	fn := func(record []string) (User, error) {
		return User{
			Name:        record[0],
			Age:         record[1],
			Description: record[2],
		}, nil
	}

	records, err := ParseCSVFile("test.csv", fn)
	if err != nil {
		t.Error(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
