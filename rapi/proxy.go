package rapi

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Proxy struct {
	router *mux.Router
}

func NewProxy() *Proxy {
	return &Proxy{
		router: mux.NewRouter(),
	}
}

func (p *Proxy) Serve() {
	http.ListenAndServe(":8080", p.router)
}

func (p *Proxy) Request(ep *Endpoint, w http.ResponseWriter, r *http.Request) {
	r.URL.Host = "localhost:8094"
	r.URL.Scheme = "http"
	r.URL.Path = "/api" + r.URL.Path

	tr := http.DefaultTransport

	resp, respErr := tr.RoundTrip(r)
	if respErr == nil {
		defer resp.Body.Close()
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	// fmt.Fprint(w, "Protected!!!!\n"+r.URL.Query().Get("querystring1"))
	// w.Write([]byte("Gorilla!\n"))
}
