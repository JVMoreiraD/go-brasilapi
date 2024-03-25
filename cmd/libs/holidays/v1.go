package holidays

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/cmd/models"
	"github.com/JVMoreiraD/go-brasilapi/cmd/utils"
)

func GetHolidays(input int) (holydays []models.Holiday, err error) {
	v1url := fmt.Sprint("feriados/v1/", input)

	var response []models.Holiday

	res := utils.HttpReq(v1url)

	if res.StatusCode == 404 {
		return []models.Holiday{}, errors.New("not found")
	}
	if res.StatusCode == 500 {
		return []models.Holiday{}, errors.New("not found")
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
