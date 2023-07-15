package mainpackage_test

import (
	"testing"

	treamie "treamie/src/package1"
	// "github.com/zbednarke/treamie/src/package1"
)

func TestFunctionToAdd(t *testing.T) {
	result := treamie.FunctionToAdd(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("FunctionToAdd(2, 3) returned %d, expected %d", result, expected)
	}
}
