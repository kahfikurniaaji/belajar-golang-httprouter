package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, err := fmt.Fprint(writer, "Hello HttpRouter")
		if err != nil {
			panic(err)
		}
		//panic("Ups")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
