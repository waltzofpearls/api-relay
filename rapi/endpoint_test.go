package rapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/relay-api/rapi"
)

func TestEndpoint(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	conf := rapi.NewConfig()
	conf.Backend.Address = strings.TrimPrefix(ts.URL, "http://")

	api := rapi.New(conf)
	require.NotNil(t, api)

	ep := rapi.NewEndpoint(api, "GET", "/foo")
	assert.NotNil(t, ep)

	req, err := http.NewRequest("GET", "/foo", nil)
	require.Nil(t, err)
	require.NotNil(t, req)

	resp := httptest.NewRecorder()
	require.NotNil(t, resp)

	ep.ServeHTTP(resp, req)
}
