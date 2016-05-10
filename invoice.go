package main

import "github.com/waltzofpearls/api-relay/rapi"

type InvoiceExternal struct {
	Total   int `json:"total"`
	Records []struct {
		Id      string       `json:"id"`
		Number  string       `json:"number"`
		Created rapi.APIDate `json:"created"`
	} `json:"records"`
}

type InvoiceInternal struct {
	Total   int `json:"total"`
	Records []struct {
		Id      string       `json:"id"`
		Number  string       `json:"number"`
		Created rapi.APIDate `json:"created"`
	} `json:"records"`
}
