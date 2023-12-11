package builtins

import (
	"fmt"
	"os"
)

// Rm removes the specified file or directory.
func Rm(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: rm <file/directory>")
		return
	}

	err := os.RemoveAll(args[0])
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed:", args[0])
	}
}