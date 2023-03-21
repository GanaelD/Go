package sandbox

import (
	"fmt"
)

func ArrayToSlice(arr []int) []int {
	sli := arr[:]
	fmt.Println(sli)
}
