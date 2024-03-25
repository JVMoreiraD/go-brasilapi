package cnpj

import (
	"encoding/json"
	"errors"
	"go-brasilapi/cmd/models"
	"go-brasilapi/cmd/utils"
	"io"
	"log"
)

func GetCNPJ(input string) (cnpj models.CNPJ, err error) {
	var v1url = "cnpj/v1/" + input

	var response models.CNPJ

	res := utils.HttpReq(v1url)
	if res.StatusCode == 404 {
		return models.CNPJ{}, errors.New("not found")
	}

	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)

	if readErr != nil {
		log.Print(err.Error())
	}

	err = json.Unmarshal((body), &response)

	if err != nil {
		log.Print(err)
	}

	return response, nil
}