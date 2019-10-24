package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var tableUsers = "users"

const AWS_REGION = "us-east-2"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func GetUsers() ([]User, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableUsers),
	}
	result, err := db.Scan(input)

	if err != nil {
		return []User{}, err
	}
	if len(result.Items) == 0 {
		return []User{}, nil
	}

	var users []User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func GetUser(id string) (User, error) {

	user := User{}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableUsers),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := db.GetItem(input)
	if len(result.Item) == 0 {
		return user, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	return user, err
}
