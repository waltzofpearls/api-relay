package main

import (
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		r.URL.Host = "localhost:8094"
		r.URL.Scheme = "http"

		tr := http.DefaultTransport

		resp, respErr := tr.RoundTrip(r)
		if respErr == nil {
			defer resp.Body.Close()
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)

		// fmt.Fprint(w, "Protected!!!!\n" + r.URL.Query().Get("querystring1"))
	})

	http.ListenAndServe(":8080", router)
}
