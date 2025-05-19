package core

// Info about an Entity that will be replicated over network
type EntityInfo struct {
	Position Vector3
	Id       uint32
}

func CreateInfo(ent *Entity) EntityInfo {
	entity := EntityInfo{}
	entity.Id = ent.id
	entity.Position = ent.position
	return entity
}
