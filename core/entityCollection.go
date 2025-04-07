package core

type EntityCollection struct {
	objects []Entity
	id      uint32 // Default is INVALID_ENTITY_ID
}

func (collection *EntityCollection) AddEntity(pos Vector3, bound CollisionPrimitive) Entity {
	entity := Entity{}
	entity.position = pos
	entity.CollisionPrimitive = bound
	collection.id = collection.id + 1
	entity.id = collection.id

	collection.objects = append(collection.objects, entity)
	return entity
}

func (collection *EntityCollection) AddEntityWithVelocity(pos Vector3, vel Vector3, bound CollisionPrimitive) Entity {
	collection.id = collection.id + 1
	entity := Entity{pos, vel, bound, collection.id}
	collection.objects = append(collection.objects, entity)
	return entity
}

// Returns the number of game objects in the collection
func (collection EntityCollection) Count() int {
	return len(collection.objects)
}
