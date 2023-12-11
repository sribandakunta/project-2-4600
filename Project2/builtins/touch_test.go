// touch_test.go

package builtins_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestTouch(t *testing.T) {
	t.Run("Create file", func(t *testing.T) {
		// Prepare a temporary directory
		tmpdir := t.TempDir()

		// Run the touch command
		builtins.Touch([]string{filepath.Join(tmpdir, "testfile.txt")})

		// Verify that the file has been created
		_, err := os.Stat(filepath.Join(tmpdir, "testfile.txt"))
		if err != nil {
			t.Errorf("Error verifying file creation: %v", err)
		}
	})

	t.Run("Error on missing file name", func(t *testing.T) {
		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the touch command with no arguments
		builtins.Touch([]string{})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Usage: touch <file>\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})

	t.Run("Error on existing file", func(t *testing.T) {
		// Prepare a temporary file
		tmpfile, err := ioutil.TempFile("", "testfile.txt")
		if err != nil {
			t.Fatalf("Error creating temporary file: %v", err)
		}
		defer os.Remove(tmpfile.Name())

		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the touch command with an existing file
		builtins.Touch([]string{tmpfile.Name()})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Error: file already exists\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})
}
