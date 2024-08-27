package tests

import (
	"fmt"
	"strings"
	"testing"
)

// T is a type used to represent a test case and manage test state.
type T struct {
	failed bool
	output func(string, ...interface{})
}

// Errorf reports a test failure and formats the error message.
func (t *T) Errorf(format string, args ...interface{}) {
	t.failed = true
	if t.output != nil {
		t.output(format+"\n", args...)
	} else {
		fmt.Printf(format+"\n", args...)
	}
}

// Fail marks the test as having failed.
func (t *T) Fail() {
	t.failed = true
}

// Failed reports whether the test has failed.
func (t *T) Failed() bool {
	return t.failed
}

// Function to test
func Add(a, b int) int {
	return a + b
}

// Test function using the standard testing.T
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got a stupid result: %d", result)
	}
}

// Test function to check error handling in test framework
func TestErrorf(t *testing.T) {
	var output strings.Builder

	// Create a custom T with an output function to capture output
	tt := &T{
		output: func(format string, args ...interface{}) {
			fmt.Fprintf(&output, format, args...)
		},
	}

	tt.Errorf("Expected %d but got %d", 5, 3)

	if output.String() != "Expected 5 but got 3\n" {
		t.Errorf("Expected 'Expected 5 but got 3\\n' but got %q", output.String())
	}

	if !tt.Failed() {
		t.Error("Expected test to be marked as failed, but it was not")
	}
}
