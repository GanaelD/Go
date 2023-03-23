package main

import (
	"example.com/sandbox"
)

func main() {
	text := sandbox.ReadText()
	words := sandbox.WordsInString(text)
	sandbox.PrintInOrder(words)
}
