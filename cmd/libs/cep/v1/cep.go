package v1

import (
	"encoding/json"
	"errors"
	"go-brasilapi/cmd/models"
	"go-brasilapi/cmd/utils"
	"io"
	"log"
)

func GetCEP(input string) (cep models.CEP, err error) {
	var v1url = "cep/v1/" + input

	var response models.CEP

	res := utils.HttpReq(v1url)

	if res.StatusCode == 404 {
		return models.CEP{}, errors.New("not found")
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
