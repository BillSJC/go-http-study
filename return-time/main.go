package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ResponseData struct {
	Error int
	Msg   string
	Data  interface{}
}

type dateNow struct {
	Timestamp int64 `json:"timestamp"`
	Year      int   `json:"year"`
	Month     int   `json:"month"`
	Day       int   `json:"day"`
}

type dateNow2 struct {
	SchoolYear string
	Semester   string
	Username   string
	dateNow
}

func main() {
	//FileReader
	//FileWriter
	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		defer func(writer http.ResponseWriter) {
			if err := recover(); err != nil {
				writer.WriteHeader(500)
				fmt.Fprintln(writer, "something went wrong")
			}
		}(writer)
		if request.Method != "POST" {
			writer.WriteHeader(405)
			return
		}
		if request.MultipartForm.Value["username"][0] == "" {
			writer.WriteHeader(401)
			return
		}
		t := time.Now()
		d := new(dateNow2)
		d.Timestamp = getTime()
		d.Year = t.Year()
		//Mon Jan 2 15:04:05 -0700 MST 2006
		m := t.Format("1")
		if mi, err := strconv.Atoi(m); err != nil {
			panic(err)
		} else {
			d.Month = mi
		}

		d.Day = t.Day()

		d.SchoolYear = "2018-2019"
		d.Semester = "2"

		resp := new(ResponseData)
		resp.Error = 0
		resp.Msg = "success"
		resp.Data = d

		b, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(writer, string(b))
		//fmt.Fprintf(writer,"%d",getTime())
	})
	//http://127.0.0.1:8888
	_ = http.ListenAndServe(":8888", nil)

}

func getTime() int64 {
	return time.Now().Unix()
}
