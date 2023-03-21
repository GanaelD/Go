package main

import (
	"fmt"

	"example.com/sandbox"
)

func main() {

	path := "C:/Users/ganae/Documents/Software engineering/Go/Go cmd commands.txt"

	text := sandbox.WordsInTextFile(path)

	fmt.Println(text)
}
