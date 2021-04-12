package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}

func main() {
	addr := ":8003"
	fmt.Printf("going to serve %s\n", addr)
	http.HandleFunc("/", hello)
	if err:= http.ListenAndServe(addr, nil); err != nil {
	fmt.Println(err)
	}
}
