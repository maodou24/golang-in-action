package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	"strings"
)

const (
	pagePadding float64 = 22
	pageBottom  float64 = 22
	textPadding float64 = 4

	// fonts
	fontNotoSansSC       = "NotoSansSC"
	fontNotoSansSemiBold = "NotoSansSCSemiBold"
	// font size
	fontSizeTitle  = 14
	fontSizeCell   = 6
	fontSizeFooter = 8

	// footer
	footnote = "github.com/maodou24"
)

type TableColumn struct {
	Title string
	Width float64
}

type TableRow []string
type Table struct {
	headerRow TableRow
	pdf       *gopdf.GoPdf

	pageRowStart int
	pageRowEnd   int
}

/*
NewTable
表格说明：
 1. 每页包含header和footer
 2. 每列可以指定宽度
 3. 文本超过列宽时自动换行
    1）列超过当前页面部分自动写到下一页
    2）未超过的列当前页面居中，下一样空白
 4. 单元格文本左对齐且垂直居中
*/
func NewTable(title string, headers []TableColumn) (*Table, error) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	err := pdf.AddTTFFont(fontNotoSansSC, "./fonts/NotoSansSC-Regular.ttf")
	if err != nil {
		return nil, err
	}

	err = pdf.AddTTFFont(fontNotoSansSemiBold, "./fonts/NotoSansSC-SemiBold.ttf")
	if err != nil {
		return nil, err
	}

	headerRow := make([]string, len(headers))
	for i := range headers {
		headerRow[i] = headers[i].Title
	}

	return &Table{
		pdf:       pdf,
		headerRow: headerRow,
	}, nil
}

func (t *Table) DrawRow(row ...TableRow) error {

}

func (t *Table) drawRow(row TableRow) error {

}

func (t *Table) addPage() error {
	err := t.addFooter()
	if err != nil {
		return err
	}
	t.pdf.AddPage()

	t.addHeader()

}

func (t *Table) addHeader() error {

}

func (t *Table) addFooter() error {
	if t.pageRowStart == t.pageRowEnd {
		return nil
	}

	err := t.pdf.SetFont(fontNotoSansSC, "", fontSizeFooter)
	if err != nil {
		return err
	}

	w, err := t.pdf.MeasureTextWidth(footnote)
	if err != nil {
		return err
	}

	t.pdf.SetXY(gopdf.PageSizeA4.W-pagePadding-w, gopdf.PageSizeA4.H-pageBottom-18)
	rect := &gopdf.Rect{W: w, H: 10}
	err = t.pdf.CellWithOption(rect, footnote, gopdf.CellOption{Align: gopdf.Middle})
	if err != nil {
		return err
	}

	pageInfo := fmt.Sprintf("%v-%v", t.pageRowStart, t.pageRowEnd)
	t.pdf.SetXY(pagePadding, gopdf.PageSizeA4.H-pageBottom-18)
	rect = &gopdf.Rect{W: w, H: 10}
	err = t.pdf.CellWithOption(rect, footnote, gopdf.CellOption{Align: gopdf.Middle})
	if err != nil {
		return err
	}
}

func (t *Table) WritePdf(filename string) error {
	err := t.addFooter()
	if err != nil {
		return err
	}
	return t.pdf.WritePdf(filename)
}
