package models

type Locals struct {
	Id    int    `json:"id"`
	Name  string `json:"nome"`
	State string `json:"estado"`
}

type Condition struct {
	CodeICA             string `json:"codigo_ica"`
	UpdatedAt           string `json:"atualizado_em"`
	AtmosphericPressure string `json:"pressao_atmosferica"`
	Visibility          string `json:"visibilidade"`
	Wind                int    `json:"vento"`
	WindDirection       int    `json:"direcao_vento"`
	Humidity            int    `json:"umidade"`
	Condition           string `json:"condicao"`
	ConditionDesc       string `json:"condicao_Desc"`
	Temperature         int    `json:"temp"`
}

type Climate struct {
	Date          string `json:"data"`
	Condition     string `json:"condicao"`
	Min           string `json:"min"`
	Max           string `json:"max"`
	UVIndex       string `json:"indice_uv"`
	ConditionDesc string `json:"condicao_desc"`
}
