package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
	"github.com/viant/endly"
	"github.com/viant/toolbox"
	"os"
	"path"
	"testing"
)

func TestClient(t *testing.T) {
	context := endly.New().NewContext(nil)
	err := setClient(context, map[string]interface{}{
		"Credentials": "4234234dasdasde",
	})
	assert.NotNil(t, err)
	_, err = getClient(context)
	assert.NotNil(t, err)
	if !toolbox.FileExists(path.Join(os.Getenv("HOME"), ".secret/aws.json")) {
		return
	}

	err = setClient(context, map[string]interface{}{
		"Credentials": "aws",
	})
	assert.Nil(t, err)
	client, err := getClient(context)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	_, ok := client.(*sqs.SQS)
	assert.True(t, ok)
}
