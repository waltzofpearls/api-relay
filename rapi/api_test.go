package rapi_test

import (
	"testing"

	"github.com/waltzofpearls/relay-api/rapi"
)

func TestCreateApi(t *testing.T) {
	config := rapi.NewConfig()
	rapi.New(config.Item)
}
