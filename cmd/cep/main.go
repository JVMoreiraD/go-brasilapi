package cep

import (
	"encoding/json"

	"io"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/models"
	"github.com/JVMoreiraD/go-brasilapi/pkg/utils"
)

func GetCEP(input string) (cep models.CEP, err error) {
	var v1url = "cep/v1/" + input

	var response models.CEP

	res, err := utils.HttpReq(v1url)

	if err != nil {
		return models.CEP{}, err
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
