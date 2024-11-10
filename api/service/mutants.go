package service

import (
	"Challenge/api/models"
	"fmt"
)



type MutantService struct {
	repo models.MutantRepository
}

func NewMutantService(mutant models.MutantRepository) *MutantService {
	return &MutantService{repo: mutant}
}

func (m MutantService) IsMutant(adns []string) (bool,error) {
	if len(adns) <= 3 {
		return false , fmt.Errorf("More DNA is required for analysis")
	}
	stats , err := m.repo.GetMutantStats()
	if err != nil {
		return false , err
	}
	var evidence int
	checkRowAndColumn(adns ,&evidence)

	if evidence <= 1 {
		checkDiagonalMatches(adns , &evidence)
	}

	m.repo.SaveMutantStatsInDynamo(stats , evidence > 1)
	return evidence > 1 , nil
}

func checkDiagonalMatches(adns []string , evidence *int) {
	rows := len(adns)
	cols := len(adns[0])
	mutantDnaBase := map[rune]bool{'A': true, 'T': true, 'C': true, 'G': true}

	for i := 0; i < rows-3; i++ {
		for j := 0; j < cols-3; j++ {
			if !mutantDnaBase[rune(adns[i][j])]{
				continue
			}
			if adns[i][j] == adns[i+1][j+1] && adns[i][j] == adns[i+2][j+2] &&
				adns[i][j] == adns[i+3][j+3] {
				*evidence++
			}
			if *evidence > 1 {
				return
			}
		}
	}

	if *evidence > 1 {
		return
	}

	for i := 3; i < rows; i++ {
		for j := 0; j < cols-3; j++ {
			if !mutantDnaBase[rune(adns[i][j])]{
				continue
			}
			if adns[i][j] == adns[i-1][j+1] && adns[i][j] == adns[i-2][j+2] &&
				adns[i][j] == adns[i-3][j+3] {
				*evidence++
			}
			if *evidence > 1 {
				return
			}
		}
	}
	return
}

func checkRowAndColumn(adns []string , evidence *int) {
	mutantDnaBase := map[rune]bool{'A': true, 'T': true, 'C': true, 'G': true}
	for i := len(adns[0]) - 1; i >= 0; i-- {
		rowTest := false
		for j := len(adns[i]) - 1; j >= 3; j-- {

			if !mutantDnaBase[rune(adns[i][j])]{
				continue
			}

			if !rowTest && adns[i][j] == adns[i][j-1] && adns[i][j-1] == adns[i][j-2] &&
				adns[i][j-2] == adns[i][j-3] {
				*evidence++
				rowTest = true
			}
		
			if adns[j][i] == adns[j-1][i] && adns[j-1][i] == adns[j-2][i] &&
				adns[j-2][i] == adns[j-3][i] {
				*evidence++
			}
		}
		if *evidence > 1 {
			return
		}

	}
	return
}

func (m MutantService) GetStatsMutant() (models.MutantStats, error) {
	return m.repo.GetMutantStats()
}
