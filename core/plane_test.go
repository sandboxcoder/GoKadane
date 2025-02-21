package core

import (
	"testing"
)

func TestBuildPlane(t *testing.T) {
	// Create a plane that lies flat on the origin (0,0,0)
	normal := Vector3{0.0, 1.0, 0.0}
	dist := -1.0
	expected := Plane{Normal: normal, Dist: dist}
	a := Vector3{-1, 1, 0}
	b := Vector3{1, 1, 0}
	c := Vector3{0, 1, -1}
	result := BuildPlane(a, b, c)
	if result != expected {
		t.Errorf("The result %+v doesnt match what was expected %+v", result, expected)
	}
}

func TestRayPlane(t *testing.T) {
	p := Plane{Normal: Vector3{0.0, 1.0, 0.0}, Dist: 0}
	rayOrigin := Vector3{0.0, 5.0, 0.0}
	rayDir := Vector3{0, -1, 0}
	var hitTime float64
	var hitPt Vector3
	result := p.RayIntersection(rayOrigin, rayDir, &hitTime, &hitPt)
	if result != true {
		t.Errorf("Expected the ray origin: %+v dir %+v to hit the plane %+v", rayOrigin, rayDir, p)
	}

	expectedTime := 5.0
	if hitTime != expectedTime {
		t.Errorf("Ray did not hit the plane at the expected time %f. Hit time was %f", expectedTime, hitTime)
	}

	expectedHitPt := Vector3{0, 0, 0}
	if hitPt != expectedHitPt {
		t.Errorf("Ray was expected to hit the plane at %+v", expectedHitPt)
	}
}

func TestClasifyPoint_OnPlane(t *testing.T) {
	plane := Plane{Normal: Vector3{0.0, 1.0, 0.0}, Dist: 0}
	pt := Vector3{0, 0, 0}
	side := plane.ClassifyPoint(pt)
	if side != HalfSpace_ON_PLANE {
		t.Errorf("Expected the point %+v to lie on the plane %+v. Not %+v", pt, plane, side)
	}
}

func TestRay(t *testing.T) {
	plane := Plane{Normal: Vector3{0.0, 1.0, 0.0}, Dist: 0}
	rayOrigin := Vector3{0, 5, 0}
	rayDir := Vector3{0, -1, 0}
	var time float64
	var hitPoint Vector3
	hitPlane := plane.RayIntersection(rayOrigin, rayDir, &time, &hitPoint)
	if hitPlane != true {
		t.Errorf("Expected to hit the plane %+v from origin %+v dir %+v", plane, rayOrigin, rayDir)
	}
}
