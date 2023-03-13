package main

import (
    "net/http"
)

func main() {

	mux := http.NewServeMux()

    mux.HandleFunc("/hello", hello)
    mux.HandleFunc("/headers", headers)

	srv := &http.Server{
		Addr: 	":8090",
		Handler: 	mux,
	}

    srv.ListenAndServe()
}