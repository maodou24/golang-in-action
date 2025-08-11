package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/signintech/gopdf"
)

const (
	pageW       float64 = 595
	pageH       float64 = 842
	pagePadding float64 = 24
	pageBottom  float64 = 16
	textPadding float64 = 4

	// header
	headerY      float64 = 90
	headerTitleH float64 = 28
	headerDateH  float64 = 14
	headerDateY  float64 = 56
	// logo image
	imgX float64 = 479
	imgW float64 = 92
	imgH float64 = 42

	// cell
	cellTextH float64 = 10
	maxCellY  float64 = pageH - 50

	// fonts
	fontNotoSansSC       = "NotoSansSC"
	fontNotoSansSemiBold = "NotoSansSCSemiBold"
	// font size
	fontSizeTitle  float64 = 20
	fontSizeDate   float64 = 10
	fontSizeFooter float64 = 8
	fontSizeCell   float64 = 6

	// footer
	footnote = "github.com/maodou24"
	footerH  = 12
	footerW  = 547
)

type TableColumn struct {
	Title string
	Width float64
}

type TableRow []string
type Table struct {
	title     string
	startTime string
	headers   []TableColumn
	headerRow TableRow
	columnX   []float64

	imgHolder gopdf.ImageHolder
	pdf       *gopdf.GoPdf

	pageRowStart int
	pageCount    int
	total        int
}

/*
NewTable
表格说明：
 1. 每页包含header和footer
 2. 每列可以指定宽度
 3. 文本超过列宽时自动换行
    1）列超过当前页面部分自动写到下一页
    2）未超过的列当前页面居中，下一样空白
 4. 单元格文本左对齐且垂直居中左对齐
*/
func NewTable(title string, total int, headers []TableColumn) (*Table, error) {
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

	imgData, err := os.ReadFile("./fonts/logo.jpg")
	if err != nil {
		return nil, err
	}

	imgHolder, err := gopdf.ImageHolderByBytes(imgData)
	if err != nil {
		return nil, err
	}

	headerRow := make([]string, len(headers))
	columnX := make([]float64, len(headers))
	for i := range headers {
		headerRow[i] = headers[i].Title
		if i == 0 {
			columnX[i] = pagePadding
		} else {
			columnX[i] = headers[i-1].Width + columnX[i-1]
		}
	}

	table := &Table{
		pdf:       pdf,
		imgHolder: imgHolder,
		headers:   headers,
		title:     title,
		columnX:   columnX,
		total:     total,
		startTime: time.Now().Format(time.DateTime),
		headerRow: headerRow,
	}

	pdf.AddPage()
	table.pageCount = 0

	err = table.addHeader()
	if err != nil {
		return nil, err
	}
	return table, nil
}

func (t *Table) DrawRow(rows ...TableRow) error {
	err := t.pdf.SetFont(fontNotoSansSC, "", fontSizeCell)
	if err != nil {
		return err
	}
	for i := range rows {
		moveToNextPage := t.pdf.GetY()+cellTextH > maxCellY
		if !moveToNextPage {
			t.pageCount++
		}

		err = t.drawRow(rows[i])
		if err != nil {
			return err
		}

		if moveToNextPage {
			t.pageCount++
		}
	}
	return nil
}

func (t *Table) drawRow(row TableRow) error {
	linesList := make([][]string, len(row))
	cellH := make([]float64, len(row))

	var maxNum int
	for i := range row {
		column := row[i]

		// if column is empty, replace empty string with -
		if len(column) == 0 {
			column = "-"
		}

		strs, err := t.pdf.SplitTextWithWordWrap(column, t.headers[i].Width-textPadding*2)
		if err != nil {
			return err
		}

		l := len(strs)
		if l > maxNum {
			maxNum = l
		}

		cellH[i] = float64(l) * cellTextH
		linesList[i] = strs
	}

	if maxNum == 0 {
		return errors.New("no data in row")
	}

	lineOffset := make([]float64, len(row))
	rowMaxH := float64(maxNum) * cellTextH
	for i := 0; i < maxNum; i++ {
		y := t.pdf.GetY()

		for j := range linesList {
			if i >= len(linesList[j]) {
				continue
			}

			// this logic is to middle the text in the cell
			// row's first line, calculate line Y offset
			if i == 0 {
				// 1. current page don't have enough space drow first line
				// 2. current page can drow all lines in row
				if y+cellTextH > maxCellY || y+rowMaxH < maxCellY {
					lineOffset[j] = (rowMaxH - cellH[j]) / 2
				} else {
					// 3. some columns exceed current page
					lineOffset[j] = max((min(rowMaxH, maxCellY-y)-cellH[j])/2, 0)
				}
			}
			if y+cellTextH > maxCellY {
				if err := t.addPage(); err != nil {
					return err
				}
				y = t.pdf.GetY()
				// if line move to next page, should follow previous page, don't need offset to middle
				if i != 0 {
					lineOffset[j] = 0
				}
			}

			t.pdf.SetXY(t.columnX[j]+textPadding, y+lineOffset[j]+textPadding)

			rect := &gopdf.Rect{W: t.headers[j].Width - 2*textPadding, H: cellTextH}
			opt := gopdf.CellOption{Align: gopdf.Middle}
			err := t.pdf.CellWithOption(rect, linesList[j][i], opt)
			if err != nil {
				return err
			}
		}

		t.pdf.SetXY(pagePadding, y+cellTextH)
	}

	t.pdf.SetXY(pagePadding, t.pdf.GetY()+2*textPadding)

	return nil
}

func (t *Table) addPage() error {
	err := t.addFooter()
	if err != nil {
		return err
	}

	t.pdf.AddPage()

	err = t.addHeader()
	if err != nil {
		return err
	}

	t.pageCount = 0
	return nil
}

func (t *Table) addHeader() error {
	err := t.pdf.SetFont(fontNotoSansSemiBold, "", fontSizeTitle)
	if err != nil {
		return err
	}

	textWidth, err := t.pdf.MeasureTextWidth(t.title)
	if err != nil {
		return err
	}

	t.pdf.SetXY(pagePadding, pagePadding)
	rect := &gopdf.Rect{W: textWidth, H: headerTitleH}
	opt := gopdf.CellOption{Align: gopdf.Middle}
	err = t.pdf.CellWithOption(rect, t.title, opt)
	if err != nil {
		return err
	}

	err = t.pdf.SetFont(fontNotoSansSC, "", fontSizeDate)
	if err != nil {
		return err
	}
	timeWidth, err := t.pdf.MeasureTextWidth(t.startTime)
	if err != nil {
		return err
	}

	t.pdf.SetXY(pagePadding, headerDateY)
	rect = &gopdf.Rect{W: timeWidth, H: headerDateH}
	opt = gopdf.CellOption{Align: gopdf.Middle}
	err = t.pdf.CellWithOption(rect, t.startTime, opt)
	if err != nil {
		return err
	}

	// logo
	err = t.pdf.ImageByHolder(t.imgHolder, imgX, pagePadding, &gopdf.Rect{W: imgW, H: imgH})
	if err != nil {
		return err
	}

	// header
	t.pdf.SetXY(pagePadding, headerY)

	err = t.pdf.SetFont(fontNotoSansSemiBold, "", fontSizeCell)
	if err != nil {
		return err
	}

	err = t.drawRow(t.headerRow)
	if err != nil {
		return err
	}

	err = t.pdf.SetFont(fontNotoSansSC, "", fontSizeCell)
	if err != nil {
		return err
	}

	t.pageRowStart += t.pageCount

	return nil
}

func (t *Table) addFooter() error {
	err := t.pdf.SetFont(fontNotoSansSC, "", fontSizeFooter)
	if err != nil {
		return err
	}

	var pageText string
	if t.pageCount > 0 {
		pageText = fmt.Sprintf("%v-%v of %v", t.pageRowStart+1, t.pageRowStart+t.pageCount, t.total)
	} else {
		pageText = fmt.Sprintf("%v-%v of %v", t.pageRowStart, t.pageRowStart, t.total)
	}
	t.pdf.SetXY(pagePadding, pageH-pageBottom-footerH)
	rect := &gopdf.Rect{W: footerW, H: footerH}
	err = t.pdf.CellWithOption(rect, pageText, gopdf.CellOption{Align: gopdf.Middle})
	if err != nil {
		return err
	}

	w, err := t.pdf.MeasureTextWidth(footnote)
	if err != nil {
		return err
	}
	t.pdf.SetXY(pageW-pagePadding-w, pageH-pageBottom-footerH)
	rect = &gopdf.Rect{W: w, H: footerH}
	err = t.pdf.CellWithOption(rect, footnote, gopdf.CellOption{Align: gopdf.Middle})
	if err != nil {
		return err
	}

	return nil
}

func (t *Table) WritePdf(filename string) error {
	err := t.addFooter()
	if err != nil {
		return err
	}
	return t.pdf.WritePdf(filename)
}
