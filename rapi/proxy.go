package rapi

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
)

type Proxy struct {
	router httptreemux.TreeMux
}

func NewProxy() *Proxy {
	return &Proxy{
		router: httptreemux.New(),
	}
}

func (p *Proxy) Serve() {
	http.ListenAndServe(":8080", p.router)
}
