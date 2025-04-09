package core

// Collection of Entities. Manages entity id assignment
type EntityCollection struct {
	objects []Entity
	id      uint32 // Default is INVALID_ENTITY_ID
}

// Attempts to add the entity. Returns true if success
func (collection *EntityCollection) addEntity(entity *Entity) bool {
	const MaxUint = ^uint32(0)
	var ret = false
	if collection.id+1 <= MaxUint {
		collection.id = collection.id + 1
		entity.id = collection.id
		collection.objects = append(collection.objects, *entity)
		ret = true
	}
	return ret
}

func (collection *EntityCollection) AddEntity(pos Vector3, bound CollisionPrimitive) Entity {
	entity := Entity{}
	entity.position = pos
	entity.CollisionPrimitive = bound
	collection.addEntity(&entity)
	return entity
}

func (collection *EntityCollection) AddEntityWithVelocity(pos Vector3, vel Vector3, bound CollisionPrimitive) Entity {
	entity := Entity{pos, vel, bound, INVALID_ENTITY_ID}
	collection.addEntity(&entity)
	return entity
}

// Returns the number of game objects in the collection
func (collection EntityCollection) Count() int {
	return len(collection.objects)
}
