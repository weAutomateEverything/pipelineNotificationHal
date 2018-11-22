package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.CloudWatchEvent) (error) {

	log.Println(request.Detail)
	return nil

}

func main() {
	lambda.Start(Handler)
}
