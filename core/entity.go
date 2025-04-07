package core

type Entity struct {
	position Vector3
	velocity Vector3
	CollisionPrimitive
	id uint32
}

const INVALID_ENTITY_ID = 0

func CreateEntity(pos Vector3, bound CollisionPrimitive) Entity {
	entity := Entity{}
	entity.position = pos
	entity.CollisionPrimitive = bound
	return entity
}

func CreateEntityWithVelocity(pos Vector3, vel Vector3, bound CollisionPrimitive) Entity {
	entity := Entity{pos, vel, bound, 0}
	return entity
}

func (entity *Entity) SetPosition(newPosition Vector3) {
	entity.position = newPosition
}

func (entity Entity) GetPosition() Vector3 {
	return entity.position
}

func (entity *Entity) SetVelocity(vel Vector3) {
	entity.velocity = vel
}

func (entity Entity) GetVelocity() Vector3 {
	return entity.velocity
}

func (entity *Entity) DoTick(dt float64) {
	newPos := entity.position.Add(entity.velocity.Mul(dt))
	entity.position = newPos
}

func (entity Entity) GetId() uint32 {
	return entity.id
}
