package main

import (
    "net/http"
)

func main() {

	mux := http.NewServeMux()

    mux.HandleFunc("/hello", hello)
    mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/", http.FileServer(http.Dir("./static")))
	srv := &http.Server{
		Addr: 	":8090",
		Handler: 	mux,
	}

    err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}