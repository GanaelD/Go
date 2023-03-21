package sandbox

// Sum sums all the elements of a slice storing integers.
func Sum(slice []int) int {
	sum := 0
	for _, nb := range slice {
		sum += nb
	}
	return sum
}

// SumEven sums all the even elements of a slice storing integers.
func SumEven(slice []int) int {
	sum := 0
	for _, nb := range slice {
		if nb%2 == 0 {
			sum += nb
		}
	}
	return sum
}

// ReverseSlice returns a new slice that is the slice passed as an argument
// with all of its elements reversed.
func ReverseSlice(slice []int) []int {
	length := len(slice)
	res := make([]int, length)
	for i := 0; i < length/2; i++ {
		res[i], res[length-(i+1)] = slice[length-(i+1)], slice[i]
	}
	// If the slice has an odd number of elements, this block is necessary to
	// obtain the element in the middle.
	if length%2 != 0 {
		res[length/2] = slice[length/2]
	}
	return res
}
