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

func TestEntityCollection_Count(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	collection := EntityCollection{}
	if collection.Count() != 0 {
		t.Errorf("result != 0")
	}
	collection.AddEntity(zeroPos, spherePrimitive)
	result := collection.Count() == 1
	if result != true {
		t.Errorf("result != expected, returned %+v instead of 1", result)
	}
	collection.AddEntity(zeroPos, spherePrimitive)

	if collection.Count() != 2 {
		t.Errorf("result != 2, returned %+v", result)
	}
}

func TestEntityCollection_Append(t *testing.T) {
	zeroPos := Vector3{0, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	collection := EntityCollection{}
	a1 := Entity{zeroPos, zeroPos, spherePrimitive, INVALID_ENTITY_ID}
	wasAdded := collection.Append(&a1)
	if collection.Count() != 1 {
		t.Errorf("Count != 1")
	}
	if wasAdded != true {
		t.Errorf("Append failed")
	}
	if a1.id != 1 {
		t.Errorf("Wrong entity id. Should have been reset")
	}
}
