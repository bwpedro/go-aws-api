package main

import (
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  if req.Path == "/patients" {
    if req.HTTPMethod == "GET" {
      return handleGetPatients(req)
    }
    if req.HTTPMethod == "POST" {
      return handleCreatePatients(req)
    }
  }

  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusMethodNotAllowed,
    Body:       http.StatusText(http.StatusMethodNotAllowed),
  }, nil
}

func main() {
  lambda.Start(router)
}

func handleGetPatients(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	patients, err := GetPatients()
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body: http.StatusText(http.StatusInternalServerError),
	  }, nil
	}
  
	js, err := json.Marshal(patients)
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body: http.StatusText(http.StatusInternalServerError),
	  }, nil
	}
  
	return events.APIGatewayProxyResponse{
	  StatusCode: http.StatusOK,
	  Body: string(js),
	}, nil
}
  
func handleCreatePatient(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var patient Patient
	err := json.Unmarshal([]byte(req.Body), &patient)
	if err != nil {
		return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
		}, nil
	}

	err = CreatePatient(patient)
	if err != nil {
		return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Created",
	}, nil
}