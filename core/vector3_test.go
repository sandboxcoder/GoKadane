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
