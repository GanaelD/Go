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

	persons, err := challenges.JsonToPerson(jsonObj)
	if err != nil {
		log.Fatalf("Error when trying to unmarshalize the json object: %v", err)
	}
	fmt.Printf("Person slice of type %T: %v\n", persons, persons)
	fmt.Println(persons[0].Name)
}
