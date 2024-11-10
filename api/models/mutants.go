package models

type Mutants struct {
	Adn []string  `json:"adn"`
}

type MutantStats struct {
    CountMutantDNA int     `json:"count_mutant_dna"` 
    CountHumanDNA  int     `json:"count_human_dna"`  
    Ratio          float64 `json:"ratio"`            
}

type MutantService interface {
	IsMutant(adns []string) (bool,error)
	GetStatsMutant() (MutantStats, error)
}

type MutantRepository interface {
	SaveMutantStatsInDynamo(stats MutantStats , isMutant bool) error 
	GetMutantStats() (MutantStats, error)
}

