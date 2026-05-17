package _encoding

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportToCsv(t *testing.T) {
	records := [][]string{
		[]string{"Name", "Age", "Description"},
		[]string{"maodou", "27", "actor"},
		[]string{"andy", "20", "to"},
	}
	err := ExportToCsv("test.csv", records)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	records, err := ReadFromCsv("test.csv")
	assert.NoError(t, err)

	for _, record := range records {
		fmt.Println(record)
	}
}
