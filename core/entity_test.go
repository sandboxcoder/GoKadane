package core

import (
	"testing"
)

func TestEntityPosition(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := CreateEntity(zeroPos, spherePrimitive)
	result := obj1.GetPosition() == zeroPos
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}

func TestEntityMoved(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := CreateEntity(zeroPos, spherePrimitive)
	newPos := Vector3{10, 11, 11}
	obj1.SetPosition(newPos)
	result := obj1.GetPosition() == newPos
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}

func TestEntityDefaultVelocity(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := CreateEntity(zeroPos, spherePrimitive)
	result := obj1.GetVelocity() == zeroPos
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}

func TestEntityVelocity(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	vel := Vector3{5, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := CreateEntityWithVelocity(zeroPos, vel, spherePrimitive)
	result := obj1.GetVelocity() == vel
	expect := true
	if result != expect {
		t.Errorf("result != expected, velocity %+v", obj1.GetVelocity())
	}
}

func TestEntity_DoTick(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	vel := Vector3{5, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := CreateEntityWithVelocity(zeroPos, vel, spherePrimitive)
	obj1.DoTick(1)
	expectedPos := vel
	result := obj1.GetPosition() == expectedPos
	expect := true
	if result != expect {
		t.Errorf("result != expected, velocity %+v", obj1.GetVelocity())
	}
}
