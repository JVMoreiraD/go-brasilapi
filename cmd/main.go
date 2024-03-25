package main

import (
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/libs/holidays"
)

func main() {
	res, err := holidays.GetHolidays(2004)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)
}
