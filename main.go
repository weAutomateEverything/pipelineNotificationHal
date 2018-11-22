package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.CloudWatchEvent) (error) {

	resp,err := request.Detail.MarshalJSON()
	if err != nil {
		return err
	}
	log.Println(string(resp))
	return nil

}

func main() {
	lambda.Start(Handler)
}
