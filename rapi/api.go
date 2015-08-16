package rapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TransformCb func() (err error)

type Api struct {
	config      *ConfigItem
	router      *mux.Router
	transformer *Transformer
}

func New(config *ConfigItem) *Api {
	a := new(Api)
	a.config = config
	a.router = mux.NewRouter()
	a.transformer = NewTransformer()
	return a
}

func (a *Api) Run() {
	http.ListenAndServe(a.config.ListenAddr, a.router)
}

func (a *Api) NewEndpoint(method, endpoint string) *Endpoint {
	return NewEndpoint(a, method, endpoint)
}
