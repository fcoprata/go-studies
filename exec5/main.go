package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Printf("valor de x: %v\n tipo de x: %T\n", x, x)
	x = 42
	fmt.Print(x)
	y := int(x)
	fmt.Printf("\nvalor de x: %v\n tipo de x: %T", y, y)
}
