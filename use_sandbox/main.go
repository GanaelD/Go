package main

import (
	"fmt"

	"example.com/sandbox"
)

func main() {
	arr := [...]int{2, 23, 435}
	slice := arr[:]

	sliceString := make([]string, 10)
	for i := 0; i < 10; i++ {
		sliceString[i] = fmt.Sprintf("%v", i)
	}

	sum := sandbox.Sum(slice)
	evenSum := sandbox.SumEven(slice)

	reversedSlice := sandbox.ReverseSlice(slice)
	reversedSliceString := sandbox.ReverseSliceF(sliceString) // equivalent to sandbox.ReverseSliceF[string](sliceString)

	toPrint := fmt.Sprintf("The sum is %v, the even sum is %v and the slice in reverse order is %v", sum, evenSum, reversedSlice)

	fmt.Println(toPrint)
	fmt.Println(reversedSliceString)
}
