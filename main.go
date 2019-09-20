package main

import (
	"fmt"
	"log"
	"net/http"
)

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Received ajax data.")

	// Invoke ParseForm before reading form values
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Acknowledge receiving the data
	fmt.Printf("User name: %s\n", r.FormValue("username"))
	fmt.Printf("Message body: %s\n", r.FormValue("message"))

	// Send back data to client
	fmt.Fprintf(w, "%s: <span> %s</span><br/>", r.FormValue("username"), r.FormValue("message"))
	
 }

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

    http.HandleFunc("/receive", receiveAjax)

	fmt.Printf("Starting HTTP server...\n")
    if err := http.ListenAndServe(":8088", nil); err != nil {
        log.Fatal(err)
    }
}