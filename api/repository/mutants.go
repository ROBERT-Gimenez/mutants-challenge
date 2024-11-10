package repository

import (
	"Challenge/api/models"
	"fmt"
	"log"
	"math"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/joho/godotenv"
)

type MutantRepository struct {
	svc *dynamodb.DynamoDB
}

func NewMutantRepository() (*MutantRepository, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	return &MutantRepository{svc: svc}, nil
}

func (r *MutantRepository) SaveMutantStatsInDynamo(stats models.MutantStats, isMutant bool) error {
		
		if isMutant {
			stats.CountMutantDNA ++
		}else{
			stats.CountHumanDNA ++
		}
		if stats.CountHumanDNA > 0 {
			stats.Ratio = math.Round(float64(stats.CountMutantDNA)/float64(stats.CountHumanDNA)*100) / 100
		} else {
			stats.Ratio = 0
		}

		input := &dynamodb.UpdateItemInput{
			TableName: aws.String("mutantChallenge"),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String("0"), 
				},
			},			
			UpdateExpression: aws.String("SET count_mutant_dna = :mutant, count_human_dna = :human, ratio = :ratio"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":mutant": {
					N: aws.String(fmt.Sprintf("%d", stats.CountMutantDNA)),
				},
				":human": {
					N: aws.String(fmt.Sprintf("%d", stats.CountHumanDNA)),
				},
				":ratio": {
					N: aws.String(fmt.Sprintf("%f", stats.Ratio)),
				},
			},
			ReturnValues: aws.String("UPDATED_NEW"), 
		}
	
		_, err := r.svc.UpdateItem(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				log.Println("Error al actualizar en DynamoDB:", aerr.Error())
			} else {
				log.Println("Error desconocido:", err)
			}
			return err
		}
	
		return nil
	}

func (r *MutantRepository) GetMutantStats() (models.MutantStats, error) {
		input := &dynamodb.GetItemInput{
			TableName: aws.String("mutantChallenge"),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String("0"),
				},
			},
		}
	
		result, err := r.svc.GetItem(input)
		if err != nil {
			log.Println("Error al obtener el elemento:", err)
			return models.MutantStats{}, err
		}
	
		if result.Item == nil {
			log.Println("No se encontró ningún elemento con id = 0")
			return models.MutantStats{}, nil
		}
	
		var stats models.MutantStats
		err = dynamodbattribute.UnmarshalMap(result.Item, &stats)
		if err != nil {
			log.Println("Error al deserializar los datos:", err)
			return models.MutantStats{}, err
		}
	
		return stats, nil
	}