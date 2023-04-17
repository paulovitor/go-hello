package main

import (
	"fmt"
	"hello/math"
)

func main() {
	resultado := math.Soma(1, 1)
	fmt.Printf("%T", resultado)
	fmt.Println("")
	fmt.Printf("%v", resultado)

	fmt.Println("")
	resultado2 := math.SomaX(1)
	fmt.Printf("%v", resultado2)
}
