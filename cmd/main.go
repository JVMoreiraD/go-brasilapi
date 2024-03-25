package main

import (
	"go-brasilapi/cmd/libs/banks"
	"log"
)

func main() {
	res, err := banks.GetBank("289")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)
}
