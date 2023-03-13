package main

import (
    "net/http"
	"log"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
    mux.HandleFunc("/hello", hello)
    mux.HandleFunc("/headers", headers)
	mux.Handle("/", fs)
	srv := &http.Server{
		Addr: 	":8090",
		Handler: 	mux,
	}

    err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}