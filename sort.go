//
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	//grab Arguments, exit with failure if none
	if len(os.Args) == 1 {
		log.Fatal("no files to sort. Please pass in a file location as an argument")
	}
	arg := os.Args[1:]
	//fmt.Println(arg)
	if arg[0] == "--help" {
		fmt.Println("This command is used to sort and join a bunch of CSV data. Accepts file name/location as argument: will take more than one and use them all together")
	}

	//create map of strings (part number without leading zeros) to int (quantity)
	m := map[string]int{}

	for _, file := range arg {
		// open CSV file to read in
		fi, err := os.Open(file)
		if err != nil {
			fmt.Println("Error reading file")
			panic(err)
		}
		defer fi.Close()

		//create reader (doesnt require a flush or a close)
		//default comma is ,
		r := csv.NewReader(fi)
		r.FieldsPerRecord = 2

		//Loop down the file
		for {
			//read next entry
			record, err := r.Read()
			//stop reading at end of file
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			partNo := record[0]
			quant, err := strconv.Atoi(record[1])
			if err != nil {
				panic(err)
			}

			m[partNo] = m[partNo] + quant
		}
	}

	//output map to file
	//open output file
	fi, err := os.Create("testdata/collapsed.csv")
	if err != nil {
		panic(err)
	}
	//close and check for returned error:
	defer fi.Close()

	//create writer:
	writer := csv.NewWriter(fi)
	defer writer.Flush()

	//increment over m
	for partNo, quant := range m {
		data := []string{partNo, strconv.Itoa(quant)}
		fmt.Println(data)
		err := writer.Write(data)
		if err != nil {
			panic(err)
		}
	}
}
