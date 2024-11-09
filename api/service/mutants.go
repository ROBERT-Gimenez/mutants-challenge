package service

import "fmt"

type MutantService struct{}

func (m MutantService) IsMutant(adns []string) bool {
	if len(adns) <= 3 {
		fmt.Printf("More DNA is required for analysis")
		return false
	}

	var evidence int
	for i := len(adns[0]) - 1; i >= 0; i-- {
		rowTest := false
		for j := len(adns[i]) - 1; j >= 3; j-- {
			if !rowTest && adns[i][j] == adns[i][j-1] && adns[i][j-1] == adns[i][j-2] && adns[i][j-2] == adns[i][j-3] {
				evidence++
				rowTest = true
			}
		
			if adns[j][i] == adns[j-1][i] && adns[j-1][i] == adns[j-2][i] && adns[j-2][i] == adns[j-3][i] {
				evidence++
			}
		}
		if evidence > 1 {
			return true
		}

	}
	if evidence <= 1 {
		checkDiagonalMatches(adns , &evidence)
	}

	return evidence > 1
}

func checkDiagonalMatches(adns []string , evidence *int) {
	rows := len(adns)
	cols := len(adns[0])

	for i := 0; i < rows-3; i++ {
		for j := 0; j < cols-3; j++ {
			if adns[i][j] == adns[i+1][j+1] &&
				adns[i][j] == adns[i+2][j+2] &&
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
			if adns[i][j] == adns[i-1][j+1] &&
				adns[i][j] == adns[i-2][j+2] &&
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
