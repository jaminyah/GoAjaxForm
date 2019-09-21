package main

import (
	"fmt"
	"log"
	"net/http"
)

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	const userErrorMsg = "Sorry, a server error occurred"
	
	fmt.Println("Received ajax data.")

	// Invoke ParseForm before reading form values
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprintf(w, userErrorMsg) // give the user a generic msg in case of a hacker
		fmt.Println(err) // log the actual error for troubleshooting
		return
	}
	for key, value := range r.Form {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Acknowledge receiving the data
	fmt.Printf("User name: %s\n", r.FormValue("username"))
	fmt.Printf("Message body: %s\n", r.FormValue("message"))

	// Send back data to client
	_, _ = fmt.Fprintf(w, "%s: <span> %s</span><br/>", r.FormValue("username"), r.FormValue("message"))
	
 }

func main() {
	const port = "8088"
	http.Handle("/", http.FileServer(http.Dir("static")))

    http.HandleFunc("/receive", receiveAjax)

	fmt.Printf("Starting HTTP server...\n")
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}
