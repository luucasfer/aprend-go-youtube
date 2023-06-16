package main

import "fmt"

var y = 20 //global - como eu nao posso utilizar := entao uso var =

func main() {
	z := 10 //local
	qualquerCoisa(z)
}

func qualquerCoisa(numRecebido int) {
	fmt.Println(numRecebido)
	fmt.Println(y)
}
