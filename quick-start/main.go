package main

import (
	"encoding/json"
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
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		data := r.PostForm.Get("data")
		m := map[string]string{"data": data}
		dataB, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		_, _ = fmt.Fprintln(w, dataB)
	})
	//listen and serve
	err := http.ListenAndServe(":8000", nil)
	fmt.Println(err)
}
