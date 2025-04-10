package core

type World struct {
	objects EntityCollection
}

func CreateWorld(collection EntityCollection) World {
	world := World{collection}
	return world
}

func CreateEmptyWorld() World {
	world := World{}
	return world
}

func (world *World) Clear() {
	world.objects = EntityCollection{}
}

func (world *World) GetCollection() EntityCollection {
	return world.objects
}

// Updates all the game objects in the game world.
// dt = deltaTime, which is the time since the game world was last updated
func (world *World) DoTick(dt float64) {
	for i := 0; i < world.objects.Count(); i++ {
		entity := world.objects.GetEntity(i)
		if entity.id != INVALID_ENTITY_ID {
			entity.DoTick(dt)
		}
	}
}

// Returns the number of game objects in the world
func (world World) GetNumEntities() int {
	return world.objects.Count()
}
