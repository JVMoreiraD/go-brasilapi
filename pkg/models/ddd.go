package models

// Estrutura do retorno do da consulta do DDD. Estado e array de cidades.
type DDD struct {
	State  string   `json:"state"`
	Cities []string `json:"cities"`
}
