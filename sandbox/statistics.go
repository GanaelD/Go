package sandbox

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// GetMean returns the mean of a slice of float64.
func GetMean(numbers []float64) float64 {
	mean := 0.
	for _, nb := range numbers {
		mean += nb
	}
	return mean / float64(len(numbers))
}

// GetMedian returns the median of a slice of float64.
func GetMedian(numbers []float64) float64 {
	length := len(numbers)
	if length == 0 {
		log.Fatalln("The slice is empty!")
	}
	sort.Float64s(numbers)
	switch length % 2 {
	// If the length of the slice is even
	case 0:
		return (numbers[length/2] + numbers[length/2+1]) / 2
	default:
		return numbers[length/2+1]
	}
}

// GetMode returns the mode of a slice of float64.
func GetMode(numbers []float64) float64 {
	// To avoid equality comparisons of float64, the keys of the map are going to be
	// the float64 parsed to strings.
	frequencies := make(map[string]int)
	for _, nb := range numbers {
		// Convert the float64 to strings using a precision of 15 decimal places to avoid rounding issues.
		key := fmt.Sprintf("%.15f", nb)
		frequencies[key] += 1
	}
	maxFrequency := 0
	mostFrequent := ""
	for key, value := range frequencies {
		if value > maxFrequency {
			mostFrequent, maxFrequency = key, value
		}
	}
	// The keys of the map being strings, we need to convert the mode back to a floating-point number.
	mostFrequentFloat, _ := strconv.ParseFloat(mostFrequent, 64)
	return mostFrequentFloat
}

// GetStatistics returns the mean, median and mode of the given slice of float64.
func GetStatistics(numbers []float64) (float64, float64, float64) {
	return GetMean(numbers), GetMedian(numbers), GetMode(numbers)
}

// PrintStatistics prints the mean, median and mode of the given slice of float64.
func PrintStatistics(numbers []float64) {
	mean, median, mode := GetStatistics(numbers)
	fmt.Printf("Here are some statistics about the slice %v:\nMean: %v\nMedian: %v\nMode: %v", numbers, mean, median, mode)
}
