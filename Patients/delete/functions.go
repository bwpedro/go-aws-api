package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var tablePatients = "patients"

const AWS_REGION = "us-east-2"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func DeletePatient(id string) error {

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tablePatients),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	if _, err := db.DeleteItem(input); err != nil {
		panic(err)
	}

	return nil
}
