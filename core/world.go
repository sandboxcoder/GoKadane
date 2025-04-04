package core

type World struct {
	objects map[int]Entity
	nextId  int
}

func CreateWorld(entities []Entity) World {
	world := CreateEmptyWorld()
	for _, entity := range entities {
		world.AddEntity(entity)
	}
	return world
}

func CreateEmptyWorld() World {
	entities := make(map[int]Entity)
	world := World{entities, 0}
	return world
}

func (world *World) ResetEntities() {
	world.objects = make(map[int]Entity)
}

func (world *World) AddEntity(entity Entity) {
	// Note: it would be good to make sure nextId < MAX_INT
	world.objects[world.nextId] = entity
	world.nextId = world.nextId + 1
}

// Updates all the game objects in the game world.
// dt = deltaTime, which is the time since the game world was last updated
func (world *World) DoTick(dt float64) {
	for _, entity := range world.objects {
		entity.DoTick(dt)
	}
}

// Returns the number of game objects in the world
func (world World) GetNumEntities() int {
	return len(world.objects)
}
