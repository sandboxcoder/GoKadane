package core

import (
	"testing"
)

func TestMaximumSubArray(t *testing.T) {
	var arr = []int{-10, -5, -1}
	result := MaximumSubArray(arr)
	expected := -1
	if result != expected {
		t.Errorf("The result %d doesnt match what was expected %d", result, expected)
	}
}
