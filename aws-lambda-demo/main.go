package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type ResponseBody struct {
	Message *string `json:"message"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	switch request.Path {
	case "/hello":
		return helloHandler(request)
	case "/hello-world-lambda":
		return bookHandler(request)
	case "/person":
		return personHandler(request)

	}
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       "Not Found",
	}, nil

}

func bookHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var book Book
	err := json.Unmarshal([]byte(request.Body), &book)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	msg := fmt.Sprintf("Book Title: %v \nBook Author: %v", book.Title, book.Author)
	responseBody := ResponseBody{&msg}

	jbytes, err := json.Marshal(responseBody)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// add the slice of bytes to the response as a string
	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}

	return res, nil
}

func personHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var person Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	msg := fmt.Sprintf("Hello %v %v", person.FirstName, person.LastName)
	responseBody := ResponseBody{&msg}

	jbytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jbytes),
	}, nil
}

func helloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Hello World")

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}
