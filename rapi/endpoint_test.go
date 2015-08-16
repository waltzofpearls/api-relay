package rapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/relay-api/rapi"
)

func TestEndpoint(t *testing.T) {

	api := rapi.New(rapi.NewConfig())
	require.NotNil(t, api)

	ep := rapi.NewEndpoint(api, "", "")
	assert.NotNil(t, ep)
}
