// Pacote focado em consultar o Cadastro de Pessoa Juridica.
package cnpj

import (
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/JVMoreiraD/go-brasilapi/pkg/models"
	"github.com/JVMoreiraD/go-brasilapi/pkg/utils"
)

// Esta função retorna uma consulta do CNPJ
func GetCNPJ(input string) (cnpj models.CNPJ, err error) {
	if input == "" {
		return models.CNPJ{}, errors.New("invalid request")
	}
	var v1url = "cnpj/v1/" + input

	var response models.CNPJ

	res, err := utils.HttpReq(v1url)
	if err != nil {
		return models.CNPJ{}, err
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
