package main

import (
	"fmt"

	"example.com/sandbox"
)

func main() {
	arr := [...]int{2, 23, 435}
	slice := arr[:]

	sum := sandbox.Sum(slice)
	evenSum := sandbox.SumEven(slice)

	reversedSlice := sandbox.ReverseSlice(slice)

	toPrint := fmt.Sprintf("The sum is %v, the even sum is %v and the slice in reverse order is %v", sum, evenSum, reversedSlice)

	fmt.Println(toPrint)
}
