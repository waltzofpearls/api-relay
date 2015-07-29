package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/", func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Protected!\n")
	})

	http.ListenAndServe(":8080", router)
}
