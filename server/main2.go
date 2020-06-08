package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
)

func http_Server(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {

	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("Received a GET request\n"))

	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		//convert reqBody from byte[] to string
		s := string(reqBody[:])

		//Test conditions with the data received
		if s == "data=data1" {
			fmt.Printf("The data is data1\n")
		}

		if s == "button=next" {
			//robotgo.KeyTap("shift", "alt")
			robotgo.KeyTap("right")
		} else if s == "button=back" {
			robotgo.KeyTap("left")
		}
		w.Write([]byte("Received a POST request\n"))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))

	}

}

func main() {
	http.HandleFunc("/", http_Server)
	http.ListenAndServe(":8000", nil)
}
