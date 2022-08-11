package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Record struct {
	Header                               bool
	DS, OS, HS, LS, CS, VS               string
	Date, Open, High, Low, Close, Volume string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	records := parse("table.csv")

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func parse(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			records = append(records, Record{
				DS:     "DATE",
				OS:     "OPEN",
				HS:     "HIGH",
				LS:     "LOW",
				CS:     "CLOSE",
				VS:     "VOLUME",
				Header: true,
			})
		} else {
			records = append(records, Record{
				Date:   row[0],
				Open:   row[1],
				High:   row[2],
				Low:    row[3],
				Close:  row[4],
				Volume: row[5],
			})
		}

	}
	return records
}
