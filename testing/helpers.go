package testing

import (
	"io/ioutil"
	"os"
)

// TempFile is a utility function which creates a file with specified content.
// The passed content can be either a string or []byte
func TempFile(content interface{}) (*os.File, func()) {
	var bytesContent []byte
	switch content := content.(type) {
	case []byte:
		bytesContent = content
	case string:
		bytesContent = []byte(content)
	default:
		panic("invalid content: it must be a string or []byte")
	}

	tmpfile, err := ioutil.TempFile("", ".tmp-file")
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		os.Remove(tmpfile.Name())
	}

	if _, err := tmpfile.Write(bytesContent); err != nil {
		cleanup()
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		cleanup()
		panic(err)
	}

	return tmpfile, cleanup
}

// TempFileFromFile is a utility function which creates a file with specified content and returns the path
func TempFileFromFile(filepath string) (*os.File, func()) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return TempFile(content)
}
