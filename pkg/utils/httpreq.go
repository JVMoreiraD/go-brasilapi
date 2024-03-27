// O pacote utils é aonde estão funções que compartilham logica entre outros pacotes
package utils

import (
	"errors"
	"log"
	"net/http"
)

// Esta função faz uma requisição HTTP para uma url composta por https://brasilapi.com.br/api/ + complemento
func HttpReq(input string) (res *http.Response, err error) {
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
	if res.StatusCode == 404 {
		return res, errors.New("not found")
	}
	if res.StatusCode == 500 {
		return res, errors.New("all services are down")
	}

	return res, nil
}
