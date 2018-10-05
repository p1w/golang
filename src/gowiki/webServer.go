package main

import (
	"fmt"
	"net/http"
)

func handler( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
