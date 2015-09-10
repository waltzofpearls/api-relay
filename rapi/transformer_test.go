package rapi_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/relay-api/rapi"
)

var same = `{
  "one": "value of one",
  "two": "value of two",
  "three": [
    {
      "four": "value of four",
      "five": "value of five"
    }
  ]
}`

type Same struct {
	One   string `json:"one"`
	Two   string `json:"two"`
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
	c, _ := v.(*To)
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
	c, _ := v.(*From)
	ct.FieldXX = c.FieldAA
	ct.FieldYY = []struct {
		FieldYY11 string `json:"field_yy_1_1"`
	}{
		{FieldYY11: c.FieldBB[0].FieldBB11},
	}
	ct.FieldZZ = struct {
		FieldZZ1 string `json:"field_zz_1"`
	}{
		FieldZZ1: c.FieldCC.FieldCC1,
	}
	return ct
}

func TestTransformSameStruct(t *testing.T) {
	tx := rapi.NewTransformer()

	out := tx.Transform([]byte(same), &Same{}, &Same{})
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

	out := tx.Transform([]byte(from), &From{}, &CustomTo{})
	require.NotNil(t, out)

	var dst bytes.Buffer
	json.Indent(&dst, out, "", "  ")

	assert.Equal(t, to, dst.String())
}

func TestTransformRequest(t *testing.T) {
	tx := rapi.NewTransformer()

	var fixture = `{"One":"this is the one", "Two":"this is the second"}`
	req, err := http.NewRequest("GET", "/test", strings.NewReader(fixture))
	require.Nil(t, err)
	require.NotNil(t, req)

	var structure struct {
		One string
		Two string
	}
	ok := tx.TransformRequest(req, structure, structure)
	require.True(t, ok)

	body, err := ioutil.ReadAll(req.Body)
	require.Nil(t, err)
	require.NotEmpty(t, body)

	assert.Equal(t, fixture, string(fixture))
}

func TestTransformResponse(t *testing.T) {
	tx := rapi.NewTransformer()

	var fixture = `{"One":"this is the one", "Two":"this is the second"}`
	res := &http.Response{
		Header:     make(http.Header),
		StatusCode: http.StatusOK,
	}
	res.Body = ioutil.NopCloser(strings.NewReader(fixture))
	require.NotNil(t, res)

	var structure struct {
		One string
		Two string
	}
	ok := tx.TransformResponse(res, structure, structure)
	require.True(t, ok)

	body, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	require.NotEmpty(t, body)

	assert.Equal(t, fixture, string(fixture))
}
