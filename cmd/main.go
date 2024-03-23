package main

import (
	v1 "go-brasilapi/cmd/libs/cep/v1"
	"log"
)

func main() {
	res, err := v1.GetCEP("60762-060")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)
}
