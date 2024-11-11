package repository

import (
	"Challenge/api/models"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBClient struct {
    mock.Mock
    UpdateItemFunc func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
    GetItemFunc    func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

func (m *MockDynamoDBClient) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
    return m.UpdateItemFunc(input)
}

func (m *MockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
    return m.GetItemFunc(input)
}

func TestSaveMutantStatsInDynamo(t *testing.T) {
    tests := []struct {
        name        string
        stats       models.MutantStats
        isMutant    bool
        mockError   error
        expectedErr error
    }{
        {
            name:     "Test mutant case",
            stats:    models.MutantStats{CountMutantDNA: 0, CountHumanDNA: 1},
            isMutant: true,
            mockError: nil,
            expectedErr: nil,
        },
        {
            name:     "Test human case",
            stats:    models.MutantStats{CountMutantDNA: 1, CountHumanDNA: 0},
            isMutant: false,
            mockError: nil,
            expectedErr: nil,
        },
        {
            name:     "Test DynamoDB update error",
            stats:    models.MutantStats{CountMutantDNA: 0, CountHumanDNA: 1},
            isMutant: true,
            mockError: fmt.Errorf("DynamoDB error"),
            expectedErr: fmt.Errorf("DynamoDB error"),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockSvc := &MockDynamoDBClient{
                UpdateItemFunc: func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
                    return nil, tt.mockError
                },
                GetItemFunc: func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
                    return &dynamodb.GetItemOutput{
                        Item: map[string]*dynamodb.AttributeValue{
                            "id": {S: aws.String("test-id")},
                        },
                    }, nil
                },
            }

            repo := &MutantRepository{
                svc: mockSvc,
            }

            err := repo.SaveMutantStatsInDynamo(tt.stats, tt.isMutant)
            if tt.expectedErr != nil {
                assert.EqualError(t, err, tt.expectedErr.Error())
                } else {
                assert.NoError(t, err)
            }
        })
    }
}

func TestSaveMutantStatsInDynamo_WhenNoHumans(t *testing.T) {
    mockSvc := &MockDynamoDBClient{
        UpdateItemFunc: func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
            return nil, nil}}
    repo := &MutantRepository{
        svc: mockSvc,
    }
	stats := models.MutantStats{
		CountMutantDNA: 10,
		CountHumanDNA:  0, 
		Ratio:          0.0,
	}

	err := repo.SaveMutantStatsInDynamo(stats, true)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, stats.Ratio)
}

func TestGetMutantStats(t *testing.T) {
	mockSvc := new(MockDynamoDBClient)

	mockSvc.GetItemFunc = func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
        return &dynamodb.GetItemOutput{
            Item: map[string]*dynamodb.AttributeValue{
                "ratio": {
                    N: aws.String("0"),
                },
                "count_mutant_dna": {
                    N: aws.String("10"),
                },
                "count_human_dna": {
                    N: aws.String("20"),
                },
            },
        }, nil
    }

    mockSvc.On("GetItem", mock.AnythingOfType("*dynamodb.GetItemInput")).Return(&dynamodb.GetItemOutput{
        Item: map[string]*dynamodb.AttributeValue{
            "ratio":            {N: aws.String("0")},
            "count_mutant_dna": {N: aws.String("10")},
            "count_human_dna":  {N: aws.String("20")},
        },
    }, nil)

	repo := MutantRepository{svc: mockSvc}

	stats, err := repo.GetMutantStats()

	assert.NoError(t, err)
	assert.Equal(t, 10, stats.CountMutantDNA)
	assert.Equal(t, 20, stats.CountHumanDNA)

}

func TestGetMutantStats_ErrorCases(t *testing.T) {
	mockSvc := new(MockDynamoDBClient)

	mockSvc.GetItemFunc = func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
		return nil, errors.New("error al obtener el elemento")
	}

	repo := MutantRepository{svc: mockSvc}

	stats, err := repo.GetMutantStats()

	assert.Error(t, err)
	assert.Equal(t, "error al obtener el elemento", err.Error())
	assert.Equal(t, models.MutantStats{}, stats)

	mockSvc.GetItemFunc = func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
		return &dynamodb.GetItemOutput{
			Item: nil,
		}, nil
	}

	stats, err = repo.GetMutantStats()

	assert.NoError(t, err)
	assert.Equal(t, models.MutantStats{}, stats)

	mockSvc.GetItemFunc = func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
		return &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String("0"),
				},
				"count_mutant_dna": {
					N: aws.String("invalid"),
				},
			},
		}, nil
	}

	stats, err = repo.GetMutantStats()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid syntax")
	assert.Equal(t, models.MutantStats{}, stats)
}