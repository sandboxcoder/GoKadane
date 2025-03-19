package core

type World struct {
	objects []Entity
}

func CreateWorld(entities []Entity) World {
	world := World{entities}
	return world
}

func CreateEmptyWorld() World {
	entities := make([]Entity, 0, 5)
	world := World{entities}
	return world
}

func (world *World) ResetEntities() {
	world.objects = make([]Entity, 0, 5)
}

func (world *World) AddEntity(entity Entity) {
	world.objects = append(world.objects, entity)
}

// Updates all the game objects in the game world.
// dt = deltaTime, which is the time since the game world was last updated
func (world *World) DoTick(dt float64) {
	for i := 0; i < len(world.objects); i++ {
		world.objects[i].DoTick(dt)
	}
}

// Returns the number of game objects in the world
func (world World) GetNumEntities() int {
	return len(world.objects)
}
