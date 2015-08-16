package rapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type Transformer struct {
}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) TransformRequest(ep *Endpoint, req *http.Request) {

}

func (t *Transformer) TransformResponse(ep *Endpoint, res *http.Response) {

}

func (t *Transformer) Transform(input []byte, v interface{}) []byte {
	err := json.Unmarshal([]byte(input), v)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %s", err)
		return nil
	}

	output, err := json.Marshal(v)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %s", err)
		return nil
	}

	return output
}
