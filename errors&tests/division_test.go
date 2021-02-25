package division

import (
	"testing"
)

func TestDivisionByZero(t *testing.T) {
	val, err := Divide(5, 0)
	if err == nil || val != 0 {
		t.Fatalf("Expected val = 0 and err = DivisionByZeroError, but was %d,%v", val, err)
	}
}

func TestDivision(t *testing.T) {
	val, err := Divide(20, 5)
	if err != nil || val != 4 {
		t.Fatalf("Expected val = 4 and err = nil, but was %d,%v", val, err)
	}
}
