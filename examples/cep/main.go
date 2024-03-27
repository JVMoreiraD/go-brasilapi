// Este pacote contem os exemplos das funções de github.com/JVMoreiraD/go-brasilapi/cep

package main

import (
	"fmt"

	"github.com/JVMoreiraD/go-brasilapi/cep"
)

func main() {
	// Retorna um cep e um erro
	cep, err := cep.GetCEP("01035-100")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cep)
}
