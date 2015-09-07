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

var from = `{
  "field_aa": "value of field aa",
  "field_bb": [
    {
      "field_bb_1_1": "value of field bb 1 1"
    }
  ],
  "field_cc": {
    "field_cc_1": "value of field cc 1"
  }
}`

var to = `{
  "field_xx": "value of field aa",
  "field_yy": [
    {
      "field_yy_1_1": "value of field b 1 1"
    }
  ],
  "field_zz": {
    "field_zz_1": "value of field c 1"
  }
}`

type From struct {
	FieldAA string
	FieldBB []struct {
		FieldBB11 string
	}
	FieldCC struct {
		FieldCC1 string
	}
}

type To struct {
	FieldXX string
	FieldYY []struct {
		FieldYY11 string
	}
	FieldZZ struct {
		FieldZZ1 string
	}
}

type CustomFrom struct {
	From
}

func (cf *CustomFrom) Transform(v interface{}) interface{} {
	v.From.FieldXX = cf.FieldAA
	return v
}

type CustomTo struct {
	To
}

func (ct *CustomTo) Transform(v interface{}) interface{} {
	return v
}

func TestTransformSameStruct(t *testing.T) {
	tx := rapi.NewTransformer()
	out := tx.Transform([]byte(same), Same{}, Same{})
	require.NotNil(t, out)

	var dst bytes.Buffer
	json.Indent(&dst, out, "", "  ")

	assert.Equal(t, same, dst.String())
}

func TestTransformCustomFrom(t *testing.T) {
	tx := rapi.NewTransformer()
	out := tx.Transform([]byte(from), CustomFrom{}, To{})
	require.NotNil(t, out)

	var dst bytes.Buffer
	json.Indent(&dst, out, "", "  ")

	assert.Equal(t, to, dst.String())
}

func TestTransformCustomTo(t *testing.T) {
	tx := rapi.NewTransformer()
	out := tx.Transform([]byte(from), From{}, CustomTo{})
	require.NotNil(t, out)

	var dst bytes.Buffer
	json.Indent(&dst, out, "", "  ")

	assert.Equal(t, to, dst.String())
}

func TestTransformRequest(t *testing.T) {
}

func TestTransformResponse(t *testing.T) {
}
