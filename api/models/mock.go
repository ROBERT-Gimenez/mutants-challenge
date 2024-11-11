package models

import "github.com/stretchr/testify/mock"

type MockMutantRepository struct {
	mock.Mock
}

func (m *MockMutantRepository) GetMutantStats() (MutantStats, error) {
	args := m.Called()
	return args.Get(0).(MutantStats), args.Error(1)
}


func (m *MockMutantRepository) SaveMutantStatsInDynamo(stats MutantStats, isMutant bool) error {
	args := m.Called(stats, isMutant)
	return args.Error(0)
}
