package core

import (
	"testing"
)

func TestSphereCollision(t *testing.T) {
	s1 := Sphere{Vector3{0, 0, 0}, 20}
	detectedCollision := checkCollision(s1, s1)
	if detectedCollision != true {
		t.Errorf("checkCollision() != expected, returned %+v instead of true", detectedCollision)
	}
}
