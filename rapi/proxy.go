package rapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Proxy struct {
	router httprouter.Router
}

func NewProxy() *Proxy {
	return &Proxy{
		router: httprouter.New(),
	}
}

func (p *Proxy) Serve() {
	http.ListenAndServe(":8080", p.router)
}
