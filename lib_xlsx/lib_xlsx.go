package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx/v3"
)

func main() {

	xlFile, err := xlsx.OpenFile("example.xlsx")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	for _, sheet := range xlFile.Sheets {
		fmt.Println("Sheet Name:", sheet.Name)
	}
}
