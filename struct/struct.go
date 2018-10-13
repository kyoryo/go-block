package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	var pointer *Vertex = &v
	pointer.X = 2
	(*pointer).Y = 1e8

	fmt.Println(v.Y)
	fmt.Println(v.X)
	fmt.Println(v)
	fmt.Println(pointer)
	fmt.Println(pointer.X == v.X)
}
