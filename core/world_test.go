package core

import (
	"testing"
)

// Creates a list of sample entities
func getEntities() []Entity {
	zeroPos := Vector3{0, 0, 0}
	vel := Vector3{1, 0, 0}
	sphere := Sphere{zeroPos, 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := CreateEntityWithVelocity(zeroPos, vel, spherePrimitive)

	pos2 := Vector3{2, 0, 0}
	sphere2 := Sphere{pos2, 5}
	obj2 := CreateEntityWithVelocity(pos2, vel, sphere2.GetPrimitive())

	list := []Entity{obj1, obj2}
	return list
}

func TestWorld_AddEntity(t *testing.T) {
	list := getEntities()
	world := World{list}
	expected := 2
	result := world.GetNumEntities()
	if result != expected {
		t.Errorf("result != expected, returned %+v instead of %+v", result, expected)
	}
}

func TestWorld_DoTick(t *testing.T) {
	list := getEntities()
	world := World{list}
	world.DoTick(1)
	expected := Vector3{1, 0, 0}
	for _, entity := range world.objects {
		vel := entity.GetVelocity()
		if vel.X != 1 {
			t.Errorf("result != expected, returned %+v instead of %+v", vel, expected)
		}
	}
}
