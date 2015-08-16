package main

type InvoiceExternal struct {
	Total   int `json:"total"`
	Records []struct {
		Id     string `json:"id"`
		Number string `json:"number"`
	} `json:"records"`
}

type InvoiceInternal struct {
	Total   int `json:"total"`
	Records []struct {
		Id     string `json:"id"`
		Number string `json:"number"`
	} `json:"records"`
}
