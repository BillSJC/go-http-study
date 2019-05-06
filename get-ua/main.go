package main

import (
	"fmt"
	"net/http"
)

/*
	use http.Request ti get something
	and now you should use recover() to ensure your service run correctly
*/

func main() {
	//	GET /
	//	return "Hello your device is " + your UA
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//use recover to save the Process
		defer func() {
			//by using recover() your service can keep up without down
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		//get your UA and return
		_, err := fmt.Fprintf(w, "Hello your device is %s", r.UserAgent())
		//panic when sth wen wrong
		if err != nil {
			panic(err)
		}
	})
	//listen anf serve
	go http.ListenAndServe("127.0.0.0:8000", nil)

}
