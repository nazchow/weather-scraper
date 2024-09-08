package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeCSV(data [][]string) {
	file, err := os.Create("weather_data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		writer.Write(value)
	}
}
