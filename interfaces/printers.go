package interfaces

import (
	"bufio"
	"fmt"
	"os"
)

type Printer interface {
	Print([]byte, error)
}

// ConsolePrinter prints a slice of data to the console using fmt.Println().
type ConsolePrinter struct{}

func (cp ConsolePrinter) Print(data []byte) error {
	fmt.Println(string(data))
	return nil
}

// FilePrinter prints a slice of data to a destination file whose path is specified in the attribute "output".
type FilePrinter struct {
	Output string
}

// Print writes data to a file, truncating its content.
func (fp FilePrinter) Print(data []byte) error {
	// Open the file for writing.
	file, err := os.Create(fp.Output)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	defer file.Close()

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	// Write the data to the file
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	fmt.Printf("Wrote data to file %v\n", fp.Output)
	return nil
}

// PrintToExistingFile adds text to an already existing text without truncating it.
func (fp FilePrinter) PrintToExistingFile(data []byte) error {
	// Read the input file
	file, err := os.OpenFile(fp.Output, os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	defer file.Close()

	// Make a buffered writer and a buffer to store what's read
	writer := bufio.NewWriter(file)

	// Write the data to the file
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	fmt.Println("Wrote data to file", fp.Output)
	return nil
}
