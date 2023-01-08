package main

import "fmt"

func main() {
	x := 22
	y := "Francisco"
	z := true
	validar_dados(x, y, z)
}

func validar_dados(x int, y string, z bool) {
	fmt.Printf("nome:%v, idade:%v, sabe python?%v", x, y, z)

}
