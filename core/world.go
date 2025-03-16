package core

type World struct {
	objects []Entity
}

func CreateWorld(entities []Entity) World {
	world := World{entities}
	return world
}

func (world *World) Init() {
	world.objects = make([]Entity, 0, 5)
}

func (world *World) AddEntity(entity Entity) {
	world.objects = append(world.objects, entity)
}

func (world *World) DoTick(dt float64) {
	for i := 0; i < len(world.objects); i++ {
		world.objects[i].DoTick(dt)
	}
}

func (world *World) GetNumEntities() int {
	return len(world.objects)
}
