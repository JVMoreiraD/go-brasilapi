package main

import (
	"go-brasilapi/cmd/libs/holidays"
	"log"
)

func main() {
	res, err := holidays.GetHolidays(2004)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(res)
}
