package testing

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

// DiffVisible is a flag which dictates whether the diff message should be shown failed assertions
var DiffVisible = true

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
	CompareValues(t, test.Actual, test.Expected)
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
	if CompareErrors(t, test.Actual[1], test.Expected[1]) {
		CompareValues(t, test.Actual[0], test.Expected[0])
	}
}

// TestErr is a Testable struct with a return value and an error
type TestErr struct {
	Name     string
	Actual   error
	Expected bool
}

// Title returns a string to display on the console when running the test
func (test TestErr) Title() string {
	return test.Name
}

// Assert is the function which gets run to compare the TestErr results
func (test TestErr) Assert(t *testing.T) {
	CompareErrors(t, test.Actual, test.Expected)
}

// CompareValues does a deep compoare of the values and types of the passed
func CompareValues(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		showDiff(actual, expected)
		if reflect.TypeOf(actual) == reflect.TypeOf(expected) {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		} else {
			t.Errorf("Expected: [%T]%v, Actual: [%T]%v", expected, expected, actual, actual)
		}
	}
}

// CompareErrors compoares actual and expected errors
func CompareErrors(t *testing.T, actualErr interface{}, expectedErr interface{}) bool {
	if (actualErr != nil) != expectedErr {
		showDiff(actualErr, expectedErr)
		t.Errorf("Expected error: %v, Actual error: %v", expectedErr, actualErr)
		return false
	}
	return true
}

// CompareValuesWithErrors compares tupples with both values and error
func CompareValuesWithErrors(t *testing.T, actual []interface{}, expected []interface{}) {
	if CompareErrors(t, actual[1], expected[1]) {
		CompareValues(t, actual[0], expected[0])
	}
}

func showDiff(actual interface{}, expected interface{}) {
	if DiffVisible {
		fmt.Println("↓↓↓↓↓↓↓↓↓↓ DIFF ↓↓↓↓↓↓↓↓↓↓")
		fmt.Println(pretty.Diff(actual, expected))
		fmt.Println("↑↑↑↑↑↑↑↑↑↑ DIFF ↑↑↑↑↑↑↑↑↑↑")
	}
}
