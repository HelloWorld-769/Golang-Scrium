package main

import (
	"testing"
)

// test function
func TestAbs(t *testing.T) {
	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1)
	}
	if Abs(0) < 0 {
		t.Error("Negative value found in abs() with", 0)
	}
	if Abs(1) < 0 {
		t.Error("Negative value found in abs() with", 1)
	}

}
