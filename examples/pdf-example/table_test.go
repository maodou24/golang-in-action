package main

import (
	"strconv"
	"testing"
)

func TestTableRowSingleLine(t *testing.T) {
	rowNum := 100

	rows := make([]TableRow, 0, rowNum)
	for i := 0; i < rowNum; i++ {
		rows = append(rows, TableRow{strconv.Itoa(i + 1)})
	}

	headers := []TableColumn{
		{Title: "ID", Width: 20},
	}

	table, err := NewTable("test", rowNum, headers)
	if err != nil {
		t.Fatal(err)
	}

	for i := range rows {
		err = table.DrawRow(rows[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	err = table.WritePdf("table.pdf")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTableRowFirstLineMorePage(t *testing.T) {
	rowNum := 400

	var str string
	for i := 0; i < rowNum; i++ {
		str += strconv.Itoa(i + 1)
	}

	headers := []TableColumn{
		{Title: "ID", Width: 20},
	}

	table, err := NewTable("test", 1, headers)
	if err != nil {
		t.Fatal(err)
	}

	err = table.DrawRow(TableRow{str})
	if err != nil {
		t.Fatal(err)
	}

	err = table.WritePdf("table.pdf")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTableRowLineMorePage(t *testing.T) {
	rowNum := 400

	var str string
	for i := 0; i < rowNum; i++ {
		str += strconv.Itoa(i + 1)
	}

	headers := []TableColumn{
		{Title: "ID", Width: 20},
	}

	table, err := NewTable("test", 3, headers)
	if err != nil {
		t.Fatal(err)
	}

	err = table.DrawRow([]TableRow{{"0"}, {str}, {"last"}}...)
	if err != nil {
		t.Fatal(err)
	}

	err = table.WritePdf("table.pdf")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTableRowToNextPage(t *testing.T) {
	rowNum := 400

	var str string
	for i := 0; i < rowNum; i++ {
		str += strconv.Itoa(i + 1)
	}

	headers := []TableColumn{
		{Title: "ID", Width: 20},
	}

	table, err := NewTable("test", 1, headers)
	if err != nil {
		t.Fatal(err)
	}

	err = table.DrawRow(TableRow{str})
	if err != nil {
		t.Fatal(err)
	}

	err = table.WritePdf("table.pdf")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTableRowToNextPageMiddle(t *testing.T) {
	rowNum := 35

	rows := make([]TableRow, 0, rowNum)
	for i := 0; i < rowNum; i++ {
		rows = append(rows, TableRow{strconv.Itoa(i + 1)})
	}

	var row string
	for i := 0; i < 30; i++ {
		row += strconv.Itoa(i + 1)
	}

	rows = append(rows, TableRow{row})

	headers := []TableColumn{
		{Title: "ID", Width: 20},
	}

	table, err := NewTable("test", len(rows), headers)
	if err != nil {
		t.Fatal(err)
	}

	err = table.DrawRow(rows...)
	if err != nil {
		t.Fatal(err)
	}

	err = table.WritePdf("table.pdf")
	if err != nil {
		t.Fatal(err)
	}
}
