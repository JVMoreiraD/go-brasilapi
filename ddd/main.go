// Este pacote propõe retornar as informações de um DDD.
package ddd

import (
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/models"
	"github.com/JVMoreiraD/go-brasilapi/pkg/utils"
)

// Esta função faz uma consulta de um DDD retornando o estado aonde ele é utilizado e as cidades que o adotam.
func GetDDD(input string) (ddd models.DDD, err error) {
	var v1url = "ddd/v1/" + input

	var response models.DDD

	res, err := utils.HttpReq(v1url)

	if err != nil {
		return models.DDD{}, errors.New("not found")
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
