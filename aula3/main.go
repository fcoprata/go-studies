package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Print(x)
	y = "10"
	fmt.Println("variavel y:", y)
	fmt.Printf("%T", z)
}
