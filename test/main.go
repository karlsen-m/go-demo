package main

import (
	"github.com/tealeg/xlsx"
)

func main() {
	xlFile, err := xlsx.OpenFile("./123123123.xlsx")
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				println(text)
			}
		}
	}

}
