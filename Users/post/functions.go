package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

var tableUsers = "users"

const AWS_REGION = "us-east-2"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func CreateUsers(user User) error {

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableUsers),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%v", uuid)),
			},
			"nome": {
				S: aws.String(user.Nome),
			},
			"sobrenome": {
				S: aws.String(user.Sobrenome),
			},
			"usuario": {
				S: aws.String(user.Usuario),
			},
			"senha": {
				S: aws.String(user.Senha),
			},
		},
	}

	_, err = db.PutItem(input)
	return err
}
