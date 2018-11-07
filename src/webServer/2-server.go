//This example shows how to handle dynamic and static (css, js etc) requests

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//handle dynamic requests
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	//static requests  e.g localhost/static/page.css (and page.css exists)
	//fs file server location
	fs := http.FileServer(http.Dir("static/"))
	//return the request page from the fs
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http listener
	http.ListenAndServe(":8080", nil)
}
