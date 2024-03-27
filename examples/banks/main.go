// Este pacote contem os exemplos das funções de github.com/JVMoreiraD/go-brasilapi/banks
package main

import (
	"log"

	"github.com/JVMoreiraD/go-brasilapi/banks"
)

func main() {
	// A função abaixo retorna todos os bancos nacionais
	bancos, _ := banks.GetAllBanks()

	log.Println(bancos)
	// Retorna um banco buscando pelo seu codigo
	banco, _ := banks.GetBank("654")

	log.Println(banco)

}
