package main

import (
	"fmt"
	"log"

	"example.com/challenges"
)

func main() {
	path := "../challenges/persons.csv"
	jsonObj, err := challenges.CsvToJson(path)
	if err != nil {
		log.Fatalf("Error when trying to get the json object: %v", err)
	}
	fmt.Println(string(jsonObj))
}
