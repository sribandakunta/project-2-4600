package builtins

import (
	"fmt"
	"io/ioutil"
)

// Cat displays the content of the specified file.
func Cat(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: cat <file>")
		return
	}

	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}