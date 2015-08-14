package rapi

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Proxy struct {
	router             *mux.Router
	listenProtocol     string
	listenAddr         string
	downstreamProtocol string
	downstreamAddr     string
}

func NewProxy() *Proxy {
	return &Proxy{
		router: mux.NewRouter(),
	}
}

func (p *Proxy) setListenProtocol(protocol string) {
	p.listenProtocol = protocol
}

func (p *Proxy) setListenAddr(addr string) {
	p.listenAddr = addr
}

func (p *Proxy) setDownstreamProtocol(protocol string) {
	p.downstreamProtocol = protocol
}

func (p *Proxy) setDownstreamAddr(addr string) {
	p.downstreamAddr = addr
}

func (p *Proxy) Serve() {
	http.ListenAndServe(":8080", p.router)
}

func (p *Proxy) Request(ep *Endpoint, w http.ResponseWriter, r *http.Request) {
	r.URL.Host = p.downstreamAddr
	r.URL.Scheme = "http"
	r.URL.Path = "/api" + ep.downstream

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
