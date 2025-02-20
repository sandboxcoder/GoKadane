package core

import (
	"testing"
)

func TestArea(t *testing.T) {
	newTriangle := Triangle{4, 6}
	area := newTriangle.Area()
	expected := 12.0
	if area != expected {
		t.Errorf("Area %f does not equals expected %f", area, expected)
	}
}
