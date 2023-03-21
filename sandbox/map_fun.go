package sandbox

import (
	"log"
	"os"
	"regexp"
)

// FindMaxValue finds the key of a map[string]int associated to the
// highest value. If the map is empty || is nil, FindMaxValue returns an empty string.
func FindMaxValue(myMap map[string]int) string {
	if len(myMap) == 0 {
		return ""
	}
	max := 0
	maxKey := ""
	for key, value := range myMap {
		if value > max {
			maxKey, max = key, value
		}
	}
	return maxKey
}

// SliceToMap converts a slice of string to a map[string]int where
// the keys are the distinct strings in the slice and the values are
// the amount of time each key appears in the slice.
func SliceToMap(slice []string) map[string]int {
	m := make(map[string]int)
	for _, str := range slice {
		m[str] += 1
	}
	return m
}

// readFile is a function used to fetch the content of a text file and put
// it inside a slice of bytes.
func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("The path %v does not exist or could not be found", path)
	}
	defer file.Close()

	// Read the content of the file.
	var content []byte
	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil || n == 0 {
			break
		}
		content = append(content, buf[:n]...)
	}
	return content
}

// TextFromFile returns a single string containing everything in the
// file whose path is passed as argument.
func TextFromFile(path string) string {
	return string(readFile(path))
}

// WordsInTextFile returns a map[string]int containing every distinct word
// in the file whose path is passed as argument and the amount of times
// they appear in the text.
func WordsInTextFile(path string) map[string]int {
	text := TextFromFile(path)
	pattern := regexp.MustCompile("[[:punct:][:space:]]+")
	words := pattern.Split(text, -1)
	return SliceToMap(words)
}

// SwapKeysAndValues takes a map as input and returns
// a new map where the keys and values of the input map
// have been swapped.
func SwapKeysAndValues[K comparable, V comparable](m *map[K]V) map[V]K {
	res := make(map[V]K)
	for key, value := range *m {
		res[value] = key
	}
	return res
}
