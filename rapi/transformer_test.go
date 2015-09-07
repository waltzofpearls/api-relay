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
      "field_yy_1_1": "value of field bb 1 1"
    }
  ],
  "field_zz": {
    "field_zz_1": "value of field cc 1"
  }
}`

type From struct {
	FieldAA string `json:"field_aa"`
	FieldBB []struct {
		FieldBB11 string `json:"field_bb_1_1"`
	} `json:"field_bb"`
	FieldCC struct {
		FieldCC1 string `json:"field_cc_1"`
	} `json:"field_cc"`
}

type To struct {
	FieldXX string `json:"field_xx"`
	FieldYY []struct {
		FieldYY11 string `json:"field_yy_1_1"`
	} `json:"field_yy"`
	FieldZZ struct {
		FieldZZ1 string `json:"field_zz_1"`
	} `json:"field_zz"`
}

type CustomFrom struct {
	From
}

func (cf *CustomFrom) Transform(v interface{}) interface{} {
	c, _ := v.(To)
	c.FieldXX = cf.FieldAA
	c.FieldYY = []struct {
		FieldYY11 string `json:"field_yy_1_1"`
	}{
		{FieldYY11: cf.FieldBB[0].FieldBB11},
	}
	c.FieldZZ = struct {
		FieldZZ1 string `json:"field_zz_1"`
	}{
		FieldZZ1: cf.FieldCC.FieldCC1,
	}
	return c
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
	out := tx.Transform([]byte(from), &CustomFrom{}, &To{})
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
