package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {

	request := events.CloudWatchEvent{}

	 err := Handler(request)

	 assert.Nil(t,err)


}
