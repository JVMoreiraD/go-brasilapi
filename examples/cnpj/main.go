// Este pacote contem os exemplos das funções de github.com/JVMoreiraD/go-brasilapi/cnpj
package main

import (
	"fmt"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/cnpj"
)

func main() {
	// Recebe um cnpj e faz retorna sua consulta
	empresaInfo, err := cnpj.GetCNPJ("")

	if err != nil {
		fmt.Println(err)
	}

	log.Println(empresaInfo)

}
