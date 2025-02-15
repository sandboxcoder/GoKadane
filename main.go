package main

import (
	"Kadane/core"
	"fmt"
)

func main() {
	fmt.Println("Simple algorithm test program")

	newTriangle := core.NewTriangle(5, 5)
	fmt.Println("area:", newTriangle.Area())

	upVector := core.Vector3{X: 0, Y: 1, Z: 0}
	rightVector := core.Vector3{X: 1, Y: 0, Z: 0}
	fmt.Println("Two othogonal vectors:", upVector.Dot(rightVector))
	fmt.Println("Mag ", upVector.Mag())
}
