package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {

	request := events.CloudWatchEvent{}
	request.Detail = []byte("{  \"pipeline\": \"pipelinenotific-Pipeline\",  \"execution-id\": \"7005a419-a525-4c50-9290-33fcaeb7e29f\",  \"stage\": \"Build\",  \"action\": \"PackageExport\",  \"state\": \"FAILED\",  \"region\": \"eu-west-1\",  \"type\": {    \"owner\": \"AWS\",    \"provider\": \"CodeBuild\",    \"category\": \"Build\",    \"version\": \"1\"  },  \"version\": 1}")

	http.HandleFunc("/api/alert/12345", func (w http.ResponseWriter, r *http.Request){
		msg,err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		s := string(msg)
		assert.Equal(t,"Code Pipeline error detected:\n" +
			"Pipeline: pipelinenotific-Pipeline\n" +
			"Stage: Build\n" +
			"Action: PackageExport",s)

	})
	go func() {
		http.ListenAndServe(":80",nil)
	}()
	os.Setenv("HAL","http://localhost:80")
	os.Setenv("GROUP","12345")
	 err := Handler(request)

	 assert.Nil(t,err)


}
