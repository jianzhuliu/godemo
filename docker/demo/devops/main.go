package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}

func main() {
	addr := ":8004"
	fmt.Printf("going to serve %s\n", addr)
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err)
	}
}
