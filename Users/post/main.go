package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleCreateUsers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var user User
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	err = CreateUsers(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers:    headers,
		Body:       "Usuário criado com sucesso!",
	}, nil
}

func main() {
	lambda.Start(handleCreateUsers)
}
