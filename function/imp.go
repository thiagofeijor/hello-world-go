package main

import (
	"fmt"
	"time"
)

func somar(a int, b int) int {
	return a + b
}

func imprimirInt(vl int) {
	fmt.Println(vl)
}

func imprimirTime(vl time.Time) {
	fmt.Println(vl)
}

func when() time.Time {
	return time.Unix(0, 0)
}
