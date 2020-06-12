/*
Moskotcho Slides Remote
Remote control slides presentations From Web browsers & any device support WiFi
----------
by Walid Amriou
https://moskotchosr.walidamriou.com
---------
*/
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/go-vgo/robotgo"
)

func main() {
	//mdns.Publish("remote.local 60 IN A 192.168.1.1")

	//To read css and js files from server
	fs := http.FileServer(http.Dir("Templates"))
	http.Handle("/Templates/", http.StripPrefix("/Templates/", fs))

	//server
	mux := pat.New()
	//if user login to 192.168.1.4:8000 from browser, will see home html page
	mux.Get("/", http.HandlerFunc(home))
	//if user send a post request
	mux.Post("/", http.HandlerFunc(send))

	//run the server
	log.Println("Listening...")
	err := http.ListenAndServe(":8000", mux)
	// if there is an error
	if err != nil {
		// log the error
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	//parse the html file index.html
	t, err := template.ParseFiles("Templates/index.html")
	// if there is an error
	if err != nil {
		// log the error
		log.Print("template parsing error: ", err)
	}
	//execute the template and pass it the HomePageVars struct to fill in the gaps
	err = t.Execute(w, nil)
	// if there is an error
	if err != nil {
		// log the error
		log.Print("template executing error: ", err)
	}
}

func send(w http.ResponseWriter, r *http.Request) {
	//Read
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", reqBody)
	//convert reqBody from byte[] to string
	s := string(reqBody[:])

	if s == "button=next" {
		robotgo.KeyTap("right")
	} else if s == "button=back" {
		robotgo.KeyTap("left")
	}
	w.Write([]byte("Received a POST request\n"))
}
