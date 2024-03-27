package models

// Estrutura dos feriados, Data, nome e tipo(Estadual ou federal)
type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Type string `json:"type"`
}
