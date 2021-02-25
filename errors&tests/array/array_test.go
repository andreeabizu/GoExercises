package array

import "testing"

var array = Arr{values: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}

func TestGetElementOverLimit(t *testing.T) {

	val, err := array.getElement(12)

	if _, ok := err.(OverTheLimitError); ok == false || val != 0 {
		t.Fatalf("Expected val=0 and err of type OverTheLimitError, but was %d,%v", val, err)
	}

}

func TestGetElementUnderLimit(t *testing.T) {

	val, err := array.getElement(-4)

	if _, ok := err.(UnderTheLimitError); ok == false || val != 0 {
		t.Fatalf("Expected val=0 and err of type UnderTheLimitError, but was %d,%v", val, err)
	}

}

func TestGetElementBetweenLimits(t *testing.T) {
	val, err := array.getElement(4)

	if err != nil || val != 5 {
		t.Fatalf("Expected val = 5 and err = nil, but was %d,%v", val, err)
	}
}

func TestGetElementAtLowerLimit(t *testing.T) {
	val, err := array.getElement(0)

	if err != nil || val != 1 {
		t.Fatalf("Expected val = 1 and err = nil, but was %d,%v", val, err)
	}
}

func TestGetElementAtUpperLimit(t *testing.T) {
	val, err := array.getElement(9)

	if err != nil || val != 10 {
		t.Fatalf("Expected val = 10 and err = nil, but was %d,%v", val, err)
	}
}
