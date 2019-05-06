package main

import (
	"fmt"
	"net/http"
)

/*
	Start a simple HTTP server and return something
*/

func main() {
	//	GET /
	//	return "Hello world!"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello world!")
		//panic when sth go wrong
		if err != nil {
			panic(err)
		}
	})
	//listen and serve
	go http.ListenAndServe("127.0.0.0:8000", nil)

}
