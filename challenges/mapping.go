package challenges

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Type Person will be used to map the csv data to JSON objects.
type Person struct {
	Name       string
	Age        int
	Occupation string
}

// ReadCSV reads a csv file and returns the data inside as a [][]string
func readCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer file.Close()

	// Create a reader to read data from the csv
	reader := csv.NewReader(file)

	// Read the data from the csv file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	return records, nil
}

// csvToStruct converts the csv data returned by readCsv and puts it inside a slice of Person.
func csvToStruct(path string) ([]Person, error) {
	// Get the data from the csv file
	records, err := readCsv(path)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	// The variable that will contain the Person
	persons := []Person{}

	// Convert the []string into Person
	for _, data := range records {
		// Try to convert the age (in column 2) to an int
		age, err := strconv.Atoi(data[1])
		if err != nil {
			fmt.Printf("One of the data contains an invalid age (column 2): %v\n", data)
			return nil, err
		}
		// Create a person using the csv data
		person := Person{Name: data[0], Age: age, Occupation: data[2]}
		persons = append(persons, person)
	}
	return persons, nil
}

// CSVToJson reads data from a csv file and returns it inside a JSON object
func CsvToJson(path string) ([]byte, error) {
	persons, err := csvToStruct(path)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	// Parse the slice of Person to a json object
	data, err := json.Marshal(persons)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return data, nil
}
