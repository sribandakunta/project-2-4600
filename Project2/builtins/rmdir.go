package builtins

import (
	"fmt"
	"os"
)

func Mkdir(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: mkdir <directory>")
		return
	}

	err := os.Mkdir(args[0], 0755)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Directory created:", args[0])
	}
}