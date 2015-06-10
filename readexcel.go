package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "testfile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("err")
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\n", cell.String())
			}
		}
	}
}
