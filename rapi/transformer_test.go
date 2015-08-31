package rapi_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/relay-api/rapi"
)

var same = `{
  "one": "value of one",
  "three": [
    {
      "five": "value of five",
      "four": "value of four"
    }
  ],
  "two": "value of two"
}`

type Same struct {
	One   string `json:"one"`
	Two   int    `json:"two"`
	Three []struct {
		Four string `json:"four"`
		Five string `json:"five"`
	} `json:"three"`
}

var customEnc = `{

}`

type CustomEnc struct {
}

var customDec = `{

}`

type CustomDec struct {
}

func TestTransformSameStruct(t *testing.T) {
	tx := rapi.NewTransformer()
	out := tx.Transform([]byte(same), Same{}, Same{})
	require.NotNil(t, out)

	var dst bytes.Buffer
	json.Indent(&dst, out, "", "  ")

	assert.Equal(t, same, dst.String())
}

func TestTransformCustomDecoder(t *testing.T) {
}

func TestTransformCustomEncoder(t *testing.T) {
}

func TestTransformRequest(t *testing.T) {
}

func TestTransformResponse(t *testing.T) {
}
