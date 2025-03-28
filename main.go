package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"log"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing request data for request %s", request.RequestContext.RequestID)

	// Process query parameters, path parameters, or request body here
	message := "Hello AWS Lambda and Netlify"

	// Example: Add the request path to the response
	if request.Path != "" {
		message += "\nRequest path: " + request.Path
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       message,
		Headers: map[string]string{
			"Content-Type":                 "text/plain",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "GET",
		},
	}, nil
}

func main() {
	log.Printf("Lambda function initialized")

	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
