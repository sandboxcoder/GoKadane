package core

import (
	"testing"
)

func TestEntityCollection_AddEntity(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	collection := EntityCollection{}
	obj1 := collection.AddEntity(zeroPos, spherePrimitive)
	result := obj1.GetId() == 1 && collection.Count() == 1
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}

func TestEntityCollection_AddEntityWithVelocity(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	vel := Vector3{5, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	collection := EntityCollection{}
	obj1 := collection.AddEntityWithVelocity(zeroPos, vel, spherePrimitive)
	result := obj1.GetVelocity() == vel && obj1.id != INVALID_ENTITY_ID && collection.Count() == 1
	expect := true
	if result != expect {
		t.Errorf("result != expected, velocity %+v", obj1.GetVelocity())
	}
}

func TestEntityCollection_GetEntity(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	collection := EntityCollection{}
	obj1 := collection.AddEntity(zeroPos, spherePrimitive)
	other := collection.Get(0)
	result := obj1.GetId() == other.GetId()
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}
