package service

import (
	"Challenge/api/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMutant(t *testing.T) {
	tests := []struct {
		name           string
		adns           []string
		mockStats      models.MutantStats
		mockError      error
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "DNA length less than 4",
			adns:           []string{"ATCG", "ATCG", "ATCG"},
			expectedResult: false,
			expectedError:  fmt.Errorf("More DNA is required for analysis"),
		},
		{
			name:           "GetMutantStats error",
			adns:           []string{"ATCG", "ATCG", "ATCG", "ATCG"},
			mockError:      fmt.Errorf("Database error"),
			expectedResult: false,
			expectedError:  fmt.Errorf("Database error"),
		},
		{
			name:           "No mutant DNA found",
			adns:           []string{"SDCG", "SDCG", "ATSD", "ATDC"},
			mockStats:      models.MutantStats{}, 
			expectedResult: false,
			expectedError:  nil,
		},
		{
			name:           "Mutant DNA found",
			adns:           []string{"ATCG", "ATCG", "DDCG", "ATCG"},
			mockStats:      models.MutantStats{},
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name: "Diagonal match (descending)",
			adns: []string{
				"GTCGG",
				"TGDGD",
				"TGGDG",
				"TGDGD",
			},
			mockStats:      models.MutantStats{},
			expectedResult: true, 
			expectedError:  nil,
		},
		{
			name: "Diagonal match (ascending)",
			adns: []string{
				"ATAGT",
				"AAGTD",
				"GGTDG",
				"GTDAA",
			},
			mockStats:      models.MutantStats{},
			expectedResult: true, 
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo := new(models.MockMutantRepository)
			mockRepo.On("GetMutantStats").Return(tt.mockStats, tt.mockError)
			mockRepo.On("SaveMutantStatsInDynamo", tt.mockStats, tt.expectedResult).Return(nil)

			service := NewMutantService(mockRepo)

			result, err := service.IsMutant(tt.adns)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}