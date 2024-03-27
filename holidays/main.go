// Este pacote consulta os feriados no Brasil.
package holidays

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/models"
	"github.com/JVMoreiraD/go-brasilapi/pkg/utils"
)

// Esta função é responsavel por retornar os feriados que temos durante um determinado ano
func GetHolidays(input int) (holydays []models.Holiday, err error) {
	v1url := fmt.Sprint("feriados/v1/", input)

	var response []models.Holiday

	res, err := utils.HttpReq(v1url)

	if err != nil {
		return []models.Holiday{}, err
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
