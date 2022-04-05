package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "hello")
	})

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	server.ListenAndServe()
}
