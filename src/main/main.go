package main

import (
	"demoProject/src/demo"
	"demoProject/src/vertex"
	"fmt"
)

func main() {
	demo.Hello()

	v := &vertex.Vertex{3, 4}

	v.Scale(0.5)
	fmt.Println(v, v.Abs())
}
