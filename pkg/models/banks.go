// O Pacote models é aonde estão armazenadas as estruturas utilizadas por outros pacotes.
package models

// Estrutura do objeto recebido pela consulta de bancos
type Bank struct {
	Ispb     string `json:"ispb"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
	FullName string `json:"fullName"`
}
