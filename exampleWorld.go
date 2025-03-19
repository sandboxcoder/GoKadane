package main

import (
	"Kadane/core"
)

func CreateGameWorld() core.World {
	world := core.CreateWorld(getEntities())
	return world
}

func getEntities() []core.Entity {
	zeroPos := core.Vector3{X: 0, Y: 0, Z: 0}
	sphere := core.Sphere{Center: zeroPos, Radius: 5}
	spherePrimitive := sphere.GetPrimitive()
	obj1 := core.CreateEntity(zeroPos, spherePrimitive)

	pos2 := core.Vector3{X: 2, Y: 0, Z: 0}
	sphere2 := core.Sphere{Center: pos2, Radius: 5}
	obj2 := core.CreateEntity(pos2, sphere2.GetPrimitive())

	list := []core.Entity{obj1, obj2}
	return list
}
