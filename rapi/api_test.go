package rapi_test

import (
	"testing"

	"github.com/waltzofpearls/api-relay/rapi"
)

func TestCreateApi(t *testing.T) {
	config := rapi.NewConfig()
	rapi.New(config)
}
