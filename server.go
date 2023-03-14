package main

import (
	"path/filepath"	
    "net/http"
	"log"
	"os"
)

var (
	StaticDir string
	LayoutPath string 
	TemplateDir string
)



func main() {

	LayoutPath = filepath.Join(os.Getenv("TEMPLATES_DIR"), "layout.html")
	StaticDir = os.Getenv("STATIC_DIR")

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(StaticDir))

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