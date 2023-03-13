package main


import (
	"fmt"
	"net/http"
	"time"
)

// func base(w http.ResponseWriter, r *http.Request) {
//
//}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("[server]: Hello handler started")
	
	defer fmt.Println("[server]: Hello handler ended")

	select {
	case <- time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("[server]: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}