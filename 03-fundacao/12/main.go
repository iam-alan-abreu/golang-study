package main

import "fmt"

func main() {

	var meuArray [3]int // declarando um array com somente 3 posições
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	println(meuArray[0])
	println(meuArray[1])
	println(meuArray[2])

	println(len(meuArray) - 1)

	// PERCORRER UM ARRAY
	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d \n", i, v)
	}
}