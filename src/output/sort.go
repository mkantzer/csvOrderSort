//
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	//grab Arguments
	arg := os.Args[1:]
	fmt.Println(arg)
	if arg[0] == "--help" {
		fmt.Println("This command is used to sort and join a bunch of CSV data. Accepts file name/location as argument: will take more than one and use them all together")
	}

	//create map of strings (part number without leading zeros) to int (quantity)
	m := map[string]int{}

	for _, file := range arg {
		// open CSV file to read in
		fi, err := os.Open(file)
		if err != nil {
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
}
