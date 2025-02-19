package core

import "testing"

func TestCross(t *testing.T) {
	up := Vector3{0, 1, 0}
	right := Vector3{1, 0, 0}
	depth := up.Cross(right)
	expected := Vector3{0, 0, -1}
	if depth != expected {
		t.Errorf("up.Cross(right) = %+v, expected %+v", depth, expected)
	}
}

func TestDot(t *testing.T) {
	up := Vector3{0, 1, 0}
	right := Vector3{1, 0, 0}
	depth := up.Dot(right)
	expected := 0.0
	if depth != expected {
		t.Errorf("up.Dot(right) = %+v, expected %+v", depth, expected)
	}
}

func TestMag(t *testing.T) {
	v := Vector3{0, 3, 0}
	expected := 3.0
	result := v.Mag()
	if v.Mag() != expected {
		t.Errorf("vector.Mag() != 3, val %f expected %f", result, expected)
	}
}

func TestAdd(t *testing.T) {
	a := Vector3{0, 3, 0}
	var b Vector3 = Vector3{1, 1, 1}
	expected := Vector3{1, 4, 1}
	result := a.Add(b)
	if result != expected {
		t.Errorf("vector.Add() != expected, val %f expected %f", result, expected)
	}
}

func TestSub(t *testing.T) {
	a := Vector3{4, 4, 4}
	var b Vector3 = Vector3{1, 1, 1}
	expected := Vector3{3, 3, 3}
	result := a.Sub(b)
	if result != expected {
		t.Errorf("vector.Sub() != expected, val %f expected %f", result, expected)
	}
}

func TestMult(t *testing.T) {
	a := Vector3{4, 4, 4}
	var b float64 = 2
	expected := Vector3{8, 8, 8}
	result := a.Mult(b)
	if result != expected {
		t.Errorf("vector.Mult() != expected, val %f expected %f", result, expected)
	}
}
