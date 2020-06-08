package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http server
	http.HandleFunc("/", http_Server)
	http.ListenAndServe(":8080", nil)

	//example of change the lang
	//robotgo.KeyTap("shift", "alt")
}

func http_Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
