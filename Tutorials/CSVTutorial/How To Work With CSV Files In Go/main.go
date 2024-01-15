package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type Record struct {
	Name   string `csv:"name"`
	Gender string `csv:"gender"`
}

type Person struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	Gender string `csv:"gender"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("data2.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV file into a slice of Record structs
	var records []Record
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	// Print the records
	for _, record := range records {
		fmt.Printf("Name: %s, Gender: %s\n", record.Name, record.Gender)
	}

	//file, err := os.Create("data2.csv")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//people := []*Person{
	//	{"Alice", 25, "Female"},
	//	{"Bob", 30, "Male"},
	//	{"Charlie", 35, "Male"},
	//}
	//
	//if err := gocsv.MarshalFile(&people, file); err != nil {
	//	panic(err)
	//}
}

//func main() {
//	// Open the CSV file
//	//file, err := os.Open("data.csv")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer file.Close()
//	//
//	//// Read the CSV data
//	//reader := csv.NewReader(file)
//	//reader.FieldsPerRecord = -1 // Allow variable number of fields
//	//data, err := reader.ReadAll()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//// Print the CSV data
//	//for _, row := range data {
//	//	for _, col := range row {
//	//		fmt.Printf("%s,", col)
//	//	}
//	//	fmt.Println()
//	//}
//
//	// Write the CSV data
//	file2, err := os.Create("data1.csv")
//	if err != nil {
//		panic(err)
//	}
//	defer file2.Close()
//
//	writer := csv.NewWriter(file2)
//	defer writer.Flush()
//	// this defines the header value and data values for the new csv file
//	headers := []string{"name", "age", "gender"}
//	data1 := [][]string{
//		{"Alice", "25", "Female"},
//		{"Bob", "30", "Male"},
//		{"Charlie", "35", "Male"},
//	}
//
//	writer.Write(headers)
//	for _, row := range data1 {
//		writer.Write(row)
//	}
//}
