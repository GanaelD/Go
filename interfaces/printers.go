package interfaces

import (
	"bufio"
	"fmt"
	"os"
)

type Printer interface {
	Print([]byte)
}

// ConsolePrinter prints a slice of data to the console using fmt.Println().
type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(data []byte) {
	fmt.Println(string(data))
}

// FilePrinter prints a slice of data to a destination file whose path is specified in the attribute "output".
type FilePrinter struct {
	Output string
}

func (fp FilePrinter) Print(data []byte) {
	// Open the file for writing.
	file, err := os.Create(fp.Output)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	// Write the data to the file
	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote data to file %v", fp.Output)
}
