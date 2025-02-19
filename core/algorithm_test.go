package core

import (
	"testing"
)

func TestMaximumSubArray_WithNegNumbers(t *testing.T) {
	var arr = []int{-10, -5, -1}
	result := MaximumSubArray(arr)
	expected := -1
	if result != expected {
		t.Errorf("The result %d doesnt match what was expected %d", result, expected)
	}
}

func TestMaximumSubArray_WithNegPosNumbers(t *testing.T) {
	var arr = []int{3, -10, -5, 1}
	result := MaximumSubArray(arr)
	expected := 3
	if result != expected {
		t.Errorf("The result %d doesnt match what was expected %d", result, expected)
	}
}

func TestMaxSubArray_OneEntry(t *testing.T) {
	arr := []int{1}
	result := MaximumSubArray(arr)
	expected := 1
	if result != expected {
		t.Errorf("The result %d doesnt match what was expected %d", result, expected)
	}
}

func TestMaxSubArray_LargestSum(t *testing.T) {
	arr := []int{-5, 8, 8}
	result := MaximumSubArray(arr)
	expected := 16
	if result != expected {
		t.Errorf("The result %d doesnt match what was expected %d", result, expected)
	}
}
