package main

import (
	"go-brasilapi/cmd/libs/cnpj"
	"log"
)

func main() {
	res, err := cnpj.GetCNPJ("")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)
}
