package main

import "fmt"

func main() {
	// a diferença entre slice e array é que o slice é dinâmico
	// e o array é estático

	s := []int{2, 4, 6, 8, 10}

	// exibindo tamanho do slice e conteudo
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	//saída: len=5 cap=5 [2 4 6 8 10]

	// dropando o slice para o tamanho 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	//saída: len=0 cap=5 []

	// dropando as duas primeiras posições do slice e exibindo o restante
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	// saída: len=3 cap=3 [6 8 10]

	// adicionando um elemento no final do slice
	// go dobra o tamanho do slice original
	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	//saída: len=6 cap=10 [2 4 6 8 10 12]
}