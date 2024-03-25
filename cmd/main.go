package main

import (
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/libs/cep"
)

func main() {
	res, err := cep.GetCEP("60762060")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)
}
