package builtins

import (
	"fmt"
	"os"
)

// Touch creates an empty file with the specified name
func Touch(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: touch <file>")
		return
	}

	file, err := os.Create(args[0])
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer file.Close()
		fmt.Println("File created:", args[0])
	}
}