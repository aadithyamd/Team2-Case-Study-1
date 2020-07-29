package CSV2JSON

import (
	"Team2CaseStudy1/pkg/Err"
	//"Team2CaseStudy1/pkg/Models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type item struct {
	Name     string
	Price    float64
	Quantity int64
}

type order struct {
	OrderID    int64
	CustomerID int64
	Restaurant string
	ItemLine   []item
	Price      float64
	Quantity   int64
	Discount   int64
	Date       string
}

// Write the orders in json file
func toJSON(orders []order, filename string) string {
	f, err := json.MarshalIndent(orders, "", "	")
	Err.CheckError(err)
	err = ioutil.WriteFile(filename+".json", f, 0644)
	Err.CheckError(err)
	fmt.Println("Output file is stored as: " + filename + ".json")
	return filename+".json"
}

// Parse the record into its parameters
func parseRecord(record []string) order {
	OID, _ := strconv.ParseInt(record[0], 10, 64)
	CID, _ := strconv.ParseInt(record[1], 10, 64)
	Rest := record[2]
	itemName := record[3]
	Price, _ := strconv.ParseFloat(record[4], 64)
	Quantity, _ := strconv.ParseInt(record[5], 10, 64)
	Discount, _ := strconv.ParseInt(record[6], 10, 64)
	date := record[7]
	itemObj := item{itemName, Price, Quantity}
	orderObj := order{OID, CID, Rest, []item{itemObj}, Price, Quantity, Discount, date}
	return orderObj
}

// Club the records in slice of order interface
func clubRecords(records [][]string) []order {
	var clubbedRecords []order
	prev := "INF"
	var orderObj order
	var itemObj item
	flag := false
	for _, record := range records {
		tempObj := parseRecord(record)
		if record[0] != prev {
			if flag == true {
				clubbedRecords = append(clubbedRecords, orderObj)
			}
			flag = true
			orderObj = tempObj
			prev = record[0]
		} else {
			itemObj = tempObj.ItemLine[0]
			orderObj.ItemLine = append(orderObj.ItemLine, itemObj)
			orderObj.Price += itemObj.Price
			orderObj.Quantity += itemObj.Quantity
		}
	}
	clubbedRecords = append(clubbedRecords, orderObj)
	fmt.Printf("Records after clubbing: %v\n", len(clubbedRecords))
	return clubbedRecords
}

// Read the csv file and store as a string in records variable
func readCSV(filename string) [][]string {
	csvFile, err := os.Open(filename)
	Err.CheckError(err)
	r := csv.NewReader(csvFile)
	_, _ = r.Read()
	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		Err.CheckError(err)
		records = append(records, record)
	}
	fmt.Printf("Records processed: %v\n", len(records))
	return records
}

// Initialise to convert the csv file to json format (json object)
func INIT(filename string) string {
	fmt.Println("Reading " + filename+".csv")
	records := readCSV(filename+".csv")
	orders := clubRecords(records)
	outputFIle := toJSON(orders, filename)
	return outputFIle
}
