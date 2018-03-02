package testing

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

// Testable is an interface that other structs can implement to aid in testing
type Testable interface {
	Title() string
	Assert(t *testing.T)
}

// Assert loops through a slice of Tesables and runs the Assert function of each of them individually
func Assert(t *testing.T, tests ...Testable) {
	for _, test := range tests {
		t.Run(test.Title(), test.Assert)
	}
}

// Tupple is a utility function to store the results of functions with multiple returns
func Tupple(a ...interface{}) []interface{} {
	return a
}

// Test is a simple Testable struct with a single return value
type Test struct {
	Name     string
	Actual   interface{}
	Expected interface{}
}

// Title returns a string to display on the console when running the test
func (test Test) Title() string {
	return test.Name
}

// Assert is the function which gets run to compare the Test results
func (test Test) Assert(t *testing.T) {
	compareValues(t, test.Actual, test.Expected)
}

// TestWithErr is a Testable struct with a return value and an error
type TestWithErr struct {
	Name     string
	Actual   []interface{}
	Expected []interface{}
}

// Title returns a string to display on the console when running the test
func (test TestWithErr) Title() string {
	return test.Name
}

// Assert is the function which gets run to compare the TestWithErr results
func (test TestWithErr) Assert(t *testing.T) {
	err := test.Actual[1]
	wantErr := test.Expected[1]
	if (err != nil) != wantErr {
		t.Errorf("Expected error: %v, Actual error: %v", wantErr, err)
		return
	}
	compareValues(t, test.Actual[0], test.Expected[0])
}

func compareValues(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		if reflect.TypeOf(actual) == reflect.TypeOf(expected) {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		} else {
			t.Errorf("Expected: [%T]%v, Actual: [%T]%v", expected, expected, actual, actual)
		}
	}
}

// TempFileWithContent is a utility function which creates a file with specified content and returns the path
func TempFileWithContent(content []byte) (f *os.File, err error) {
	tmpfile, err := ioutil.TempFile("", ".tmp-file")
	if err != nil {
		return tmpfile, err
	}

	if _, err := tmpfile.Write(content); err != nil {
		return tmpfile, err
	}
	if err := tmpfile.Close(); err != nil {
		return tmpfile, err
	}
	return tmpfile, nil
	// Note: Make sure to cleanup the file when done by calling: defer os.Remove(tmpfile.Name())
}

// TempFileFromFile is a utility function which creates a file with specified content and returns the path
func TempFileFromFile(filepath string) (f *os.File, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return &os.File{}, err
	}
	return TempFileWithContent(content)
	// Note: Make sure to cleanup the file when done by calling: defer os.Remove(tmpfile.Name())
}
