package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name myEvent) (events.APIGatewayProxyResponse, error) {

	data, err := json.Marshal(name)
	msg := fmt.Sprintf("Hello %s! :), Welcome to AWS Lambda", data[Name])

	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
