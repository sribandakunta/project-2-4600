// cp_test.go

package builtins_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestCp(t *testing.T) {
	t.Run("Copy file", func(t *testing.T) {
		// Prepare a temporary source file with content
		content := []byte("Test content for cp command.")
		tmpfile, err := ioutil.TempFile("", "sourcefile.txt")
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

		// Prepare a temporary destination file
		tmpdir := t.TempDir()
		destfile := filepath.Join(tmpdir, "destinationfile.txt")

		// Run the cp command
		builtins.Cp([]string{tmpfile.Name(), destfile})

		// Verify that the destination file has the same content
		destContent, err := ioutil.ReadFile(destfile)
		if err != nil {
			t.Fatal(err)
		}

		if string(destContent) != string(content) {
			t.Errorf("Expected destination content: %v, got: %v", string(content), string(destContent))
		}
	})

	t.Run("Error on missing source file", func(t *testing.T) {
		// Prepare a temporary destination file
		tmpdir := t.TempDir()
		destfile := filepath.Join(tmpdir, "destinationfile.txt")

		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the cp command with a non-existent source file
		builtins.Cp([]string{"nonexistentfile.txt", destfile})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Error opening source: open nonexistentfile.txt: no such file or directory\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})
}
