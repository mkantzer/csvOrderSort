package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// open output file
	fi, err := os.Create("orderList.csv")
	if err != nil {
		panic(err)
	}
	//close and check for returned error:
	defer fi.Close()

	//create writer:
	writer := csv.NewWriter(fi)
	defer writer.Flush()

	for i := 0; i < 100000; i++ {
		//generate data
		partNo := rand.Intn(1000)
		quantity := rand.Intn(10)
		//set partNo to a string, pad it with zeros to 4 digits:
		sPartNo := fmt.Sprintf("%04d", partNo)
		data := []string{sPartNo, strconv.Itoa(quantity)}
		fmt.Println(data)
		err := writer.Write(data)
		if err != nil {
			panic(err)
		}
	}
}
