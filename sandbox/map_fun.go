package sandbox

import (
	"log"
	"os"
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

// returnStringFileContent uses readFile to return a
// slice of string containing the content of the text file
func returnStringFileContent(path string) []string {
	content := readFile(path)
	contentString := make([]string, len(content))
	for i := 0; i < len(content); i++ {
		contentString[i] = string(content[i])
	}
	return contentString
}

func WordsInTextFile(path string) map[string]int {
	return SliceToMap(returnStringFileContent(path))
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
