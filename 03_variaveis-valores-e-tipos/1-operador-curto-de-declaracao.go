package main

import "fmt"

func main() {

	//TIPAGEM AUTOMATICA
	x := 10 //operador curto de declaração - só funciona dentro de um codeblock
	y := "olá, bom dia"
	a := 10.0
	fmt.Printf("X tem o valor: %v, e o tipo %T\n", x, x)
	fmt.Printf("Y tem o valor: %v, e o tipo %T\n", y, y)
	fmt.Printf("A tem o valor: %v, e o tipo %T\n", a, a)


	x = 20 //operador de atribuição
			//Não posso fazer x:=20 porque esta variavel ja existe
	fmt.Printf("X tem o valor: %v, e o tipo %T\n", x, x)


	x, z := 40, 30 //dessa forma a declaração funciona
	fmt.Printf("X tem o valor: %v, e o tipo %T\n", x, x)
	fmt.Printf("Z tem o valor: %v, e o tipo %T\n", z, z)

}
