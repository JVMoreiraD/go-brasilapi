package models

type Bank struct {
	Ispb     string `json:"ispb"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
	FullName string `json:"fullName"`
}
