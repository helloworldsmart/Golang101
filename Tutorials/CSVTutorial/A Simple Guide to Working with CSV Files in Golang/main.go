package main

import (
	"encoding/csv"
	"fmt"
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
func main() {
	file, err := os.Open("data2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	reader.Comma = '|'

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, each := range csvData {
		fmt.Println(each)
	}
}
