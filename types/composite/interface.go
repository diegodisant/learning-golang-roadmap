package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

func (square *Square) Area() int {
	return square.side * square.side
}

type Scalable interface {
	Scale(amplitude float64)
	Distance() float64
}

type Vertex struct {
	X float64
	Y float64
}

func (vertex *Vertex) Scale(amplitude float64) {
	vertex.X = vertex.X * amplitude
	vertex.Y = vertex.Y * amplitude
}

func (vertex *Vertex) Distance() float64 {
	return math.Sqrt((vertex.X * vertex.X) + (vertex.Y * vertex.Y))
}

func main() {
	var square Shape

	square = &Square{side: 4}

	fmt.Println(square.Area())

	fmt.Println(reflect.TypeOf(square))

	var vertex Scalable

	vertex = &Vertex{X: 2, Y: 2}

	fmt.Println(vertex.Distance())

	vertex.Scale(4)

	fmt.Println(vertex.Distance())

	fmt.Println(reflect.TypeOf(vertex))
}
