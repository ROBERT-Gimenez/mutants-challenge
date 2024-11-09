package models

type Mutants struct {
	Adn []string  `json:"adn"`
}

type MutantService interface {
	IsMutant(adn []string) bool
}