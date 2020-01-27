package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const PI float64 = 3.1415
	const raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println("Área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	var g byte = 255
	fmt.Println("O byte é", reflect.TypeOf(g))
}
