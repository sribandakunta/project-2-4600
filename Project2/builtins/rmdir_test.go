// rmdir_test.go

package builtins_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestMkdir(t *testing.T) {
	t.Run("Create directory", func(t *testing.T) {
		// Prepare a temporary directory
		tmpdir := filepath.Join(os.TempDir(), "testdir")

		// Run the mkdir command
		builtins.Mkdir([]string{tmpdir})

		// Verify that the directory has been created
		_, err := os.Stat(tmpdir)
		if err != nil {
			t.Errorf("Error verifying directory creation: %v", err)
		}
	})

	t.Run("Error on missing directory name", func(t *testing.T) {
		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the mkdir command with no arguments
		builtins.Mkdir([]string{})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Usage: mkdir <directory>\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})

	t.Run("Error on existing directory", func(t *testing.T) {
		// Prepare a temporary directory
		tmpdir := t.TempDir()

		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the mkdir command with an existing directory
		builtins.Mkdir([]string{tmpdir})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Error: mkdir " + tmpdir + ": file exists\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})
}
