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
		return false , fmt.Errorf("more DNA is required for analysis")
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

	for rowIndex := 0; rowIndex < rows-3; rowIndex++ {
		for j := 0; j < cols-3; j++ {
			if j+3 < cols && rowIndex+3 < rows && !mutantDnaBase[rune(adns[rowIndex][j])]{
				continue
			}
			if adns[rowIndex][j] == adns[rowIndex+1][j+1] &&
				adns[rowIndex][j] == adns[rowIndex+2][j+2] &&
				adns[rowIndex][j] == adns[rowIndex+3][j+3] {
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

	for rowIndex := 3; rowIndex < rows; rowIndex++ {
		for j := 0; j < cols-3; j++ {
			if j+3 < cols && rowIndex+3 < rows && !mutantDnaBase[rune(adns[rowIndex][j])]{
				continue
			}
			if adns[rowIndex][j] == adns[rowIndex-1][j+1] &&
				adns[rowIndex][j] == adns[rowIndex-2][j+2] &&
				adns[rowIndex][j] == adns[rowIndex-3][j+3] {
				*evidence++
			}
			if *evidence > 1 {
				return
			}
		}
	}
}

func checkRowAndColumn(adns []string, evidence *int) {
    mutantDnaBase := map[rune]bool{'A': true, 'T': true, 'C': true, 'G': true}
    rows := len(adns)
    if rows == 0 {
        return
    }
    cols := len(adns[0])

    for rowIndex := 0; rowIndex < rows; rowIndex++ {
        rowTest := false

        for j := 0; j < cols; j++ {

            if j >= 3 && mutantDnaBase[rune(adns[rowIndex][j])] {
                if !rowTest && adns[rowIndex][j] == adns[rowIndex][j-1] &&
					adns[rowIndex][j] == adns[rowIndex][j-2] &&
					adns[rowIndex][j] == adns[rowIndex][j-3] {
                    *evidence++
                    rowTest = true
                }
            }

            if rowIndex >= 3 && mutantDnaBase[rune(adns[rowIndex][j])] {
                if adns[rowIndex][j] == adns[rowIndex-1][j] &&
				 	adns[rowIndex][j] == adns[rowIndex-2][j] &&
					adns[rowIndex][j] == adns[rowIndex-3][j] {
                    *evidence++
                }
            }

            if *evidence > 1 {
                return
            }
        }
    }
}

func (m MutantService) GetStatsMutant() (models.MutantStats, error) {
	return m.repo.GetMutantStats()
}
