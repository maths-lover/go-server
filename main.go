// a simple web server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/about" {
		if r.Method != "GET" {
			http.Error(w, "Any method other than GET is not supported on this route", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "Hi! I'm Maths Lover, I am a B.E. Student and I love maths and computer")
	} else {
		http.Error(w, "404, Page not found", http.StatusNotFound)
		return
	}
}

func basicFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/basicForm" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v\n", err)
		}
		extract := r.FormValue
		fmt.Fprintln(w, "Form parse successfull")
		fmt.Fprintf(w, "Your name: %s\nYour Phone Number: %s\n", extract("name"), extract("phoneNum"))
	} else {
		http.Error(w, "404, Page not found", http.StatusNotFound)
		return
	}
}

func main() {
	server := http.FileServer(http.Dir("static"))
	http.Handle("/", server)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/basicForm", basicFormHandler)

	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
