package main

import (
	"example.com/interfaces"
)

func main() {
	fp := interfaces.FilePrinter{Output: "file.txt"}
	toPrint := "This is a test. Please remain calm."
	fp.Print([]byte(toPrint))
}
