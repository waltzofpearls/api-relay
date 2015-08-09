package rapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Proxy struct {
	router mux.Router
}

func NewProxy() *Proxy {
	return &Proxy{
		router: *mux.NewRouter(),
	}
}

func (p *Proxy) Serve() {
	http.ListenAndServe(":8080", &p.router)
}
