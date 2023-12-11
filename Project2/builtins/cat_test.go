// cat_test.go

package builtins_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestCat(t *testing.T) {
	t.Run("Display content of the file", func(t *testing.T) {
		// Prepare a temporary file with content
		content := []byte("Test content for cat command.")
		tmpfile, err := ioutil.TempFile("", "testfile.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		if _, err := tmpfile.Write(content); err != nil {
			tmpfile.Close()
			t.Fatal(err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatal(err)
		}

		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the cat command
		builtins.Cat([]string{tmpfile.Name()})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the output
		expectedOutput := string(content) + "\n"
		if string(out) != expectedOutput {
			t.Errorf("Expected output: %v, got: %v", expectedOutput, string(out))
		}
	})

	t.Run("Error on missing file", func(t *testing.T) {
		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the cat command with a non-existent file
		builtins.Cat([]string{"nonexistentfile.txt"})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Error reading file: open nonexistentfile.txt: no such file or directory\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})
}
