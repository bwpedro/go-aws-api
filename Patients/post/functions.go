package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

var tablePatients = "patients"

const AWS_REGION = "us-east-2"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func CreatePatients(patient Patient) error {

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tablePatients),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%v", uuid)),
			},
			"nome": {
				S: aws.String(patient.Nome),
			},
			"sobrenome": {
				S: aws.String(patient.Sobrenome),
			},
			"cpf": {
				S: aws.String(patient.CPF),
			},
			"rg": {
				S: aws.String(patient.RG),
			},
			"nascimento": {
				S: aws.String(patient.Nascimento),
			},
			"sexo": {
				S: aws.String(patient.Sexo),
			},
			"obs": {
				S: aws.String(patient.Obs),
			},
		},
	}

	_, err = db.PutItem(input)
	return err
}
