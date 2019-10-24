package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleUpdatePatients(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var patient Patient
	err := json.Unmarshal([]byte(req.Body), &patient)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	err = UpdatePatients(patient)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Paciente atualizado com sucesso!",
	}, nil
}

func main() {
	lambda.Start(handleUpdatePatients)
}
