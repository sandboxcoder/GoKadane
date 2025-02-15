package main

import (
	"Kadane/core"
	"fmt"
)

func main() {
	fmt.Println("Simple algorithm test program")

	t := core.NewTriangle(5, 5)
	fmt.Println(t.Area())
}
