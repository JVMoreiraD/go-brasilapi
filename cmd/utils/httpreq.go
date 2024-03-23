package utils

import (
	"log"
	"net/http"
)

func HttpReq(input string) (res *http.Response) {
	var baseURl = "https://brasilapi.com.br/api/"
	url := baseURl + input
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Print(err.Error())
	}

	res, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Print(err.Error())
	}

	return res
}
