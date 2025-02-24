package core

type Sphere struct {
	Center Vector3
	Radius float64
}

// Static Sphere vs Sphere test
func checkCollision(s1 Sphere, s2 Sphere) bool {
	distanceSquared := s1.Center.Sub(s2.Center).LengthSquared()
	radiusSum := s1.Radius + s2.Radius
	return distanceSquared <= radiusSum*radiusSum
}
