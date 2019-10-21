package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)
const AWS_REGION = "sa-east-1"
const TABLE_NAME = "go-serverless-api"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func GetPatients() ([]Patient, error) {

	input := &dynamodb.ScanInput{
	  TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)
	if err != nil {
	  return []Patient{}, err
	}
	if len(result.Items) == 0 {
	  return []Patient{}, nil
	}
  
	var patients[]Patient
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &patients)
	if err != nil {
	  return []Patient{}, err
	}
  
	return patients, nil
  }
  
  func CreatePatient(patient Patient) error {
  
	uuid, err := uuid.NewV4()
	if err != nil {
	  return err
	}

	input := &dynamodb.PutItemInput{
	  TableName: aws.String(tableName),
	  Item: map[string]*dynamodb.AttributeValue{
		"id": {
		  S: aws.String(fmt.Sprintf("%v", uuid)),
		},
		"name": {
		  S: aws.String(patient.Name),
		},
	  },
	}
  
	_, err = db.PutItem(input)
	return err
  }