package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// All go htt.Request objects have a built in context handler
	ctx := req.Context()
	fmt.Println("[server]: Hello handler started")

	// defer calls allow us to schedule a function to run when the
	// function exits
	defer fmt.Println("[server]: Hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("[server]: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	// simple loop to iterate over a both
	// keys and values of dictionary like object
	for name, headers := range req.Header {

		// for loop runs over the header content
		// lopos over a range in go return both the item and
		// the iteration count of the loop unlike python
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func serveTemplate(w http.ResponseWriter, req *http.Request) {
	fp := filepath.Join("templates", "index.html")

	tmpl, _ := template.ParseFiles(LayoutPath, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
