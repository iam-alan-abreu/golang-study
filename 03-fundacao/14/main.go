package main

import "fmt"

func main() {

	// _ blank identifier
	// _, _, _, _ = 1, 2, 3, 4

	salarios := map[string]int{"Wesley": 2000, "Alan": 3000, "Lorran": 4000}

	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Nolan"] = 5000

	fmt.Println(salarios["Nolan"])

	for nome, salario := range salarios {
		fmt.Printf("Nome: %s - Salario: %d \n", nome, salario)
	}

	// estou ignorando o nome usando o blank identifier _)
	for _, salario := range salarios {
		fmt.Printf("Salario: %d \n",  salario)
	}

}