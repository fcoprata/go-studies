package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Printf("x: %v, tipo: %T\n", x, x)
	x = 42
	fmt.Print(x)
}
