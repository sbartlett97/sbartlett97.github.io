package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	StaticDir   string
	LayoutPath  string
	TemplateDir string
)

func main() {

	LayoutPath = filepath.Join(os.Getenv("TEMPLATES_DIR"), "layout.html")
	StaticDir = os.Getenv("STATIC_DIR")

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./" + StaticDir))
	CssFs := http.FileServer(http.Dir("./css"))
	ImgFs := http.FileServer(http.Dir("./images"))
	FontsFs := http.FileServer(http.Dir("./fonts"))
	ScriptsFs := http.FileServer(http.Dir("./scripts"))
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/", serveTemplate)

	// without striping the /static/ preix for static urls, the server will treat
	// /static as a folder in the route of the filesystem
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/css/", http.StripPrefix("/css/", CssFs))
	mux.Handle("/images/", http.StripPrefix("/images/", ImgFs))
	mux.Handle("/fonts/", http.StripPrefix("/fonts/", FontsFs))
	mux.Handle("/js/", http.StripPrefix("/js/", ScriptsFs))

	srv := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
