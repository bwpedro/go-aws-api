package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var tablePatients = "patients"

const AWS_REGION = "us-east-2"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func GetPatients() ([]Patient, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(tablePatients),
	}
	result, err := db.Scan(input)
	if err != nil {
		return []Patient{}, err
	}
	if len(result.Items) == 0 {
		return []Patient{}, nil
	}

	var patients []Patient
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &patients)
	if err != nil {
		return []Patient{}, err
	}

	return patients, nil
}

func GetPatient(id string) (Patient, error) {

	patient := Patient{}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tablePatients),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := db.GetItem(input)
	if len(result.Item) == 0 {
		return patient, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &patient)
	if err != nil {
		fmt.Println(err.Error())
		return patient, err
	}

	return patient, err
}
