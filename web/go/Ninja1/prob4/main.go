/*
   Problem Statement
   - Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
   - Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.
*/

package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080",nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	record := parse("table.csv")
	
	tpl , err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Error =",err)
	}

	err = tpl.Execute(res, record)
	
	
}

func parse(file string) []Record {

	f , err := os.Open("table.csv")
	if err != nil{
		log.Fatalln("Error =",err)
	}

	defer f.Close()

	src  := csv.NewReader(f)

	rows, err := src.ReadAll()
	if err != nil {
		log.Fatalln("Error =",err)
	}

	records := make([]Record, 0 , len(rows))
	
	for i, j := range rows {
		if i == 0 {

			continue
		}

		date , _ := time.Parse("2006-01-02", j[0])
		open , _ := strconv.ParseFloat(j[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records
}
