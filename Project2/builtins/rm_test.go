// rm_test.go

package builtins_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestRm(t *testing.T) {
	t.Run("Remove file", func(t *testing.T) {
		// Prepare a temporary file
		tmpfile, err := ioutil.TempFile("", "testfile.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up

		// Run the rm command
		builtins.Rm([]string{tmpfile.Name()})

		// Verify that the file has been removed
		_, err = os.Stat(tmpfile.Name())
		if !os.IsNotExist(err) {
			t.Errorf("Expected file to be removed, but it still exists.")
		}
	})

	t.Run("Remove directory", func(t *testing.T) {
		// Prepare a temporary directory
		tmpdir, err := ioutil.TempDir("", "testdir")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpdir) // clean up

		// Run the rm command
		builtins.Rm([]string{tmpdir})

		// Verify that the directory has been removed
		_, err = os.Stat(tmpdir)
		if !os.IsNotExist(err) {
			t.Errorf("Expected directory to be removed, but it still exists.")
		}
	})

	t.Run("Error on missing file/directory", func(t *testing.T) {
		// Redirect stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Run the rm command with a non-existent file/directory
		builtins.Rm([]string{"nonexistentfile.txt"})

		// Restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read the captured output
		out, _ := ioutil.ReadAll(r)

		// Verify the error message
		expectedError := "Error: remove nonexistentfile.txt: no such file or directory\n"
		if string(out) != expectedError {
			t.Errorf("Expected error: %v, got: %v", expectedError, string(out))
		}
	})
}
