package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var tablePatients = "patients"

const AWS_REGION = "us-east-2"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func UpdatePatients(patient Patient) error {

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tablePatients),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(patient.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":nome": {
				S: aws.String(patient.Nome),
			},
			":sobrenome": {
				S: aws.String(patient.Sobrenome),
			},
			":cpf": {
				S: aws.String(patient.CPF),
			},
			":rg": {
				S: aws.String(patient.RG),
			},
			":nascimento": {
				S: aws.String(patient.Nascimento),
			},
			":sexo": {
				S: aws.String(patient.Sexo),
			},
			":obs": {
				S: aws.String(patient.Obs),
			},
		},
		UpdateExpression: aws.String("set nome = :nome, sobrenome = :sobrenome, cpf = :cpf, rg = :rg, nascimento = :nascimento, sexo = :sexo, obs = :obs"),
	}

	_, err := db.UpdateItem(input)
	return err
}
