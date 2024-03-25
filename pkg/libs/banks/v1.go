package banks

import (
	"encoding/json"
	"errors"

	"io"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/models"
	"github.com/JVMoreiraD/go-brasilapi/pkg/utils"
)

func GetBank(input string) (banks models.Bank, err error) {
	var v1url = "banks/v1/" + input
	var response models.Bank
	res, err := utils.HttpReq(v1url)
	if err != nil {
		return models.Bank{}, err
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
func GetAllBanks() (banks []models.Bank, err error) {
	var v1url = "banks/v1/"
	var response []models.Bank
	res, err := utils.HttpReq(v1url)
	if err != nil {
		return []models.Bank{}, errors.New("not found")
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
