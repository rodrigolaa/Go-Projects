package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	heightLenght float64
	widthLenght  float64
}

func main() {

	triangle := triangle{15, 5}
	square := square{4}

	printArea(triangle)
	printArea(square)

}

func (s square) getArea() float64 {

	return s.sideLenght * s.sideLenght

}

func (t triangle) getArea() float64 {

	return 0.5 * t.heightLenght * t.widthLenght

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
