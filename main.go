package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kyokomi/emoji"
	"net/http"
	"os"
	"strings"
)

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.CloudWatchEvent) (error) {

	v := event{}
	err := json.Unmarshal(request.Detail, &v)
	if err != nil {
		return err
	}

	e := ":white_check_mark:"
	switch v.State {
	case "FAILED":
		e = ":x:"
	case "SUCCEEDED":
		e = ":checkered_flag:"
		
		
	
	}
	if v.State == "FAILED" {
		
	}

	msg := emoji.Sprintf("%v Code Pipeline Event:\n"+
		"Event: %v\n"+
		"Pipeline: %v\n",e, v.State, v.Pipeline)

	resp, err := http.Post(os.Getenv("HAL"), "application/text", strings.NewReader(msg))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil

}

func main() {
	lambda.Start(Handler)
}

type event struct {
	Pipeline string `json:"pipeline"`
	Stage    string `json:"stage"`
	Action   string `json:"action"`
	State    string `json:"state"`
}
