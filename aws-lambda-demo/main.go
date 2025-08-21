package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Hello world!")

	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World",
	}

	return res, nil
}
