package main

import (
	"encoding/csv"
	"log"
	"os"
)

//func main() {
//	file, err := os.Open("data.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	reader := csv.NewReader(file)
//	reader.FieldsPerRecord = -1
//
//	csvData, err := reader.ReadAll()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, each := range csvData {
//		fmt.Println(each)
//	}
//
//	//
//	for {
//		record, err := reader.Read()
//
//		if err == io.EOF {
//			break
//		}
//
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(record)
//	}
//
//}

// MARK: Reading files with different delimiters
//func main() {
//	file, err := os.Open("data2.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	reader := csv.NewReader(file)
//	reader.FieldsPerRecord = -1
//
//	reader.Comma = '|'
//
//	csvData, err := reader.ReadAll()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, each := range csvData {
//		fmt.Println(each)
//	}
//}

// MARK: Writing CSV Files
func main() {
	file, err := os.Create("data3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{"name", "age", "city", "job", "company", "phone"},
		{"Kevin", "30", "New York", "Manager", "Folio Inc.", "355-025-0525"},
		{"John", "25", "Los Angeles", "Developer", "Folio Inc.", "555-230-2139"},
		{"Mary", "28", "Chicago", "HR", "Folio Inc.", "331-359-8311"},
		{"Elizabeth", "35", "New York", "Data Analyst", "Apple Inc.", "113-602-9009"},
	}

	for _, value := range data {
		writer.Write(value)
	}

	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}
}
