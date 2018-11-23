package testhelper

import (
	"os"
	"testing"
)

// TestSetEnv provides a way to change an environment variable and
// returns a callback to revert it back to the original value
func TestSetEnv(t *testing.T, key string, value string) func() {
	t.Helper()
	originalValue := os.Getenv(key)
	os.Setenv(key, value)
	return func() { os.Setenv(key, originalValue) }
}
