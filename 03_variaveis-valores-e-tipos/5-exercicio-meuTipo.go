package main

import "fmt"

type meuTipo int

var x meuTipo
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%T", y)

}
