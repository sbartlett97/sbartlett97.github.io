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
	mux.HandleFunc("/", serveTemplate)

	// without striping the /static/ preix for static urls, the server will treat
	// /static as a folder in the route of the filesystem
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	srv := &http.Server{
		Addr: 	":8090",
		Handler: 	mux,
	}

    err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}