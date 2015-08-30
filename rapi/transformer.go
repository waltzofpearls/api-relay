package rapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Transformable interface {
	TransformRequest(*http.Request, interface{}, interface{}) bool
	TransformResponse(*http.Response, interface{}, interface{}) bool
}

type Customizable interface {
	Transform(v interface{}) *interface{}
}

type Transformer struct {
}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) TransformRequest(r *http.Request, ex, in interface{}) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		return false
	}

	out := t.Transform(body, ex, in)
	if out == nil {
		return false
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(out))
	r.Header.Del("Content-Length")

	return true
}

func (t *Transformer) TransformResponse(r *http.Response, in, ex interface{}) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return false
	}

	out := t.Transform(body, in, ex)
	if out == nil {
		return false
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(out))
	r.Header.Del("Content-Length")

	return true
}

func (t *Transformer) Transform(body []byte, dec, enc interface{}) []byte {
	err := json.Unmarshal(body, &dec)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %s", err)
		return nil
	}

	if c, ok := dec.(Customizable); ok {
		enc = c.Transform(&enc)
	} else if c, ok := enc.(Customizable); ok {
		enc = c.Transform(&dec)
	} else {
		enc = &dec
	}

	out, err := json.Marshal(enc)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %s", err)
		return nil
	}

	return out
}
