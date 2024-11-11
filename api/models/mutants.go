package models

import "github.com/aws/aws-sdk-go/service/dynamodb"

type DynamoDBClient interface {
    UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}
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

