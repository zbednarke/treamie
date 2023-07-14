package mainpackage_test

import (
	"testing"
)

func TestFunctionToAdd(t *testing.T) {
	result := mainpackage.FunctionToAdd(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("FunctionToAdd(2, 3) returned %d, expected %d", result, expected)
	}
}
