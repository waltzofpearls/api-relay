package rapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Transformable interface {
	TransformRequest(*http.Request, interface{}) bool
	TransformResponse(*http.Response, interface{}) bool
}

type Transformer struct {
}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) TransformRequest(req *http.Request, v interface{}) bool {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		return false
	}

	out := t.Transform(body, v)
	if out == nil {
		return false
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(out))
	req.Header.Del("Content-Length")

	return true
}

func (t *Transformer) TransformResponse(res *http.Response, v interface{}) bool {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return false
	}

	out := t.Transform(body, v)
	if out == nil {
		return false
	}

	res.Body = ioutil.NopCloser(bytes.NewBuffer(out))
	res.Header.Del("Content-Length")

	return true
}

func (t *Transformer) Transform(in []byte, v interface{}) []byte {
	err := json.Unmarshal([]byte(in), v)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %s", err)
		return nil
	}

	out, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %s", err)
		return nil
	}

	return out
}
