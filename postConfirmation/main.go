package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	fmt.Printf("PostConfirmation for UserAttributes: %s\n", event.Request.UserAttributes)
	fmt.Printf("PostConfirmation for user: %s\n", event.UserName)
	fmt.Printf("PostConfirmation for Request: %s\n", event.Request)
	fmt.Printf("PostConfirmation event: %v\n", event)

	return event, nil
}

func main() {
	lambda.Start(Handler)
}
