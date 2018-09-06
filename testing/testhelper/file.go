package testhelper

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// TestFile is a utility function which creates a file with specified content.
// The passed content can be either a string or []byte
func TestFile(t *testing.T, content interface{}) (*os.File, func()) {
	t.Helper()

	var bytesContent []byte
	switch content := content.(type) {
	case []byte:
		bytesContent = content
	case string:
		bytesContent = []byte(content)
	default:
		t.Fatal("invalid content: it must be a string or []byte")
	}

	tmpfile, err := ioutil.TempFile("", ".tmp-file")
	if err != nil {
		t.Fatal(err)
	}
	cleanup := func() {
		os.Remove(tmpfile.Name())
	}

	if _, err := tmpfile.Write(bytesContent); err != nil {
		cleanup()
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		cleanup()
		t.Fatal(err)
	}

	return tmpfile, cleanup
}

// TestFileWithExtension is a utility function which creates a file with specified content
// and extension. The passed content can be either a string or []byte
func TestFileWithExtension(t *testing.T, content interface{}, extension string) (*os.File, func()) {
	t.Helper()

	tmpfile, cleanup := TestFile(t, content)
	newFilename := fmt.Sprintf("%v.%v", tmpfile.Name(), extension)
	if err := os.Rename(tmpfile.Name(), newFilename); err != nil {
		cleanup()
		t.Fatal(err)
	}
	cleanup = func() {
		os.Remove(newFilename)
	}
	tmpfile, err := os.Open(newFilename)
	if err != nil {
		cleanup()
		t.Fatal(err)
	}
	return tmpfile, cleanup
}

// TestFileFromFile is a utility function which creates a file with the
// content of the file read from the specified filepath
func TestFileFromFile(t *testing.T, filepath string) (*os.File, func()) {
	t.Helper()

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatal(err)
	}
	return TestFile(t, content)
}
