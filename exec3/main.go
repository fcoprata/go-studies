package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x, y, z = 42, "james bond", true

	s := fmt.Sprint(x, y, z)
	fmt.Print(s)
}
