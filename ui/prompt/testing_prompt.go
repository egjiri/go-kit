package prompt

import (
	"bufio"
	"io"
	"testing"
)

// TestSetReader is a test helper function which allows the user to
// overwrite the default reader and provides a callback in the return
// to be able to restore it back to the original using defer
func TestSetReader(t *testing.T, r io.Reader) func() {
	t.Helper()
	defaultReader := reader
	reader = bufio.NewReader(r)
	return func() { reader = defaultReader }
}
