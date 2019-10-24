package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleDeletePatients(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	err := DeletePatient(req.QueryStringParameters["id"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Paciente deletado com sucesso!",
	}, nil
}

func main() {
	lambda.Start(handleDeletePatients)
}
