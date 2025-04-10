package main

import (
	"Kadane/core"
)

func CreateGameWorld() core.World {
	world := core.CreateWorld(getEntities())
	return world
}

func getEntities() core.EntityCollection {
	collection := core.EntityCollection{}
	zeroPos := core.Vector3{X: 0, Y: 0, Z: 0}
	sphere := core.Sphere{Center: zeroPos, Radius: 5}
	spherePrimitive := sphere.GetPrimitive()
	collection.AddEntity(zeroPos, spherePrimitive)

	pos2 := core.Vector3{X: 2, Y: 0, Z: 0}
	sphere2 := core.Sphere{Center: pos2, Radius: 5}
	collection.AddEntity(pos2, sphere2.GetPrimitive())

	return collection
}
