package main

import (
	"fmt"
	"net/http"
)

func main() {

	//the webserver in oneline. "/"  is the url to listen for. the rest is always the same
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//the response from the server. 
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	//listen for traffic on port 8080.  nil is the default router
	http.ListenAndServe(":8080", nil)
}
