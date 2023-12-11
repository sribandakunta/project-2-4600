package builtins

import (
	"fmt"
	"io"
	"os"
)

// Cp copies the source file or directory to the destination.
func Cp(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: cp <source> <destination>")
		return
	}

	source, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error opening source:", err)
		return
	}
	defer source.Close()

	destination, err := os.Create(args[1])
	if err != nil {
		fmt.Println("Error creating destination:", err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println("Error copying:", err)
	} else {
		fmt.Println("Copied:", args[0], "to", args[1])
	}
}