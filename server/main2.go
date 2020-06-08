package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		s := string(reqBody[:])

		//data_receive := string([]byte{reqBody})
		if s == "data=data1" {
			fmt.Printf("The data is data1\n")
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
