package sandbox

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetInputMap gets a list of persons (their name and age) from the user
// and returns it inside a map[string]int where the keys are the names
// of the persons and the values their age.
func GetInputInMap() map[string]int {
	res := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt the user to enter the name and age of the person formated as such: first_name last_name age
		fmt.Print("Please enter the name and age of the person (first_name last_name age): ")
		if !scanner.Scan() {
			break
		}
		// Get the answer of the user
		input := scanner.Text()
		fields := strings.Fields(input)
		// If the length of the answer is != 3 then the input is invalid
		if len(fields) != 3 {
			fmt.Println("Invalid input. Please enter name and age separated by spaces.")
			continue
		}
		firstName, lastName, ageStr := fields[0], fields[1], fields[2]
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Invalid age. Please enter a valid number.")
			continue
		}
		name := fmt.Sprintf("%v %v", firstName, lastName)
		res[name] = age
		// Ask the user if there are more persons to store or not. If the user doesn't answer y or Y, the program stops,
		// else it starts a new iteration of the loop.
		fmt.Print("Do you have more persons to give? y/n ")
		if !scanner.Scan() || strings.ToLower(scanner.Text()) != "y" {
			break
		}
	}
	return res
}

// ReadText reads a text input from the user and returns a map where the keys are the
// distinct words and the values the number of times they appear in the text.
func ReadText() map[string]int {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your text:")
	if !scanner.Scan() {
		log.Fatalf("Invalid text. Exiting the function.")
	}
	text := scanner.Text()
	// This is to make no difference between capital and non capital letters.
	text = strings.ToLower(text)
	// The distinct words and the number of times they appear in the string.
	return WordsInString(text)
}
