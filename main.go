package main

import (
	"Kadane/core"
	"fmt"
)

func main() {
	fmt.Println("Simple algorithm test program")

	newTriangle := core.Triangle{Base: 5, Height: 5}
	fmt.Println("area:", newTriangle.Area())

	upVector := core.Vector3{X: 0, Y: 1, Z: 0}
	leftVector := core.Vector3{X: 1, Y: 0, Z: 0}
	fmt.Println("Two orthogonal vectors:", upVector.Dot(leftVector))
	fmt.Println("Magnitude of the up vector:", upVector.Mag())

	testVec := core.Vector3{X: 0, Y: 2, Z: 0}
	testVec.Normalize()
	fmt.Println("Normalized vec ", testVec)

	crossProduct := upVector.Cross(leftVector)
	fmt.Println("Cross product (up X left) = back vector:", crossProduct)

	arr := []int{-10, -3, -1}
	maxSub := core.MaximumSubArray(arr)
	fmt.Println("maximumSubArray value for {-10,-3,-1} is", maxSub)
}
