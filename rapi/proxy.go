package rapi

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Proxy struct {
	config *Config
	router *mux.Router
	tr     http.RoundTripper
}

func NewProxy(c *Config) *Proxy {
	return &Proxy{
		config: c,
		router: mux.NewRouter(),
	}
}

func (p *Proxy) Serve() {
	http.ListenAndServe(p.config.item.ListenAddr, p.router)
}

func (p *Proxy) Request(ep *Endpoint, w http.ResponseWriter, r *http.Request) {
	if p.tr == nil {
		p.tr = http.DefaultTransport
	}

	r.URL.Host = p.config.item.Downstream
	r.URL.Scheme = "http"
	r.URL.Path = "/api" + ep.internalPath

	res, resErr := p.tr.RoundTrip(r)
	if resErr != nil {
		log.Printf("Error requesting downstream %s: %s", r.URL.Path, resErr)
	} else {
		defer res.Body.Close()
	}

	w.WriteHeader(res.StatusCode)
	_, ioErr := io.Copy(w, res.Body)

	if ioErr != nil {
		log.Printf("Error writting response: %s", ioErr)
	}

	// fmt.Fprint(w, "Protected!!!!\n"+r.URL.Query().Get("querystring1"))
	// w.Write([]byte("Gorilla!\n"))
}
