package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleCreatePatients(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var patient Patient
	err := json.Unmarshal([]byte(req.Body), &patient)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	err = CreatePatients(patient)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Paciente criado com sucesso!",
	}, nil
}

func main() {
	lambda.Start(handleCreatePatients)
}
