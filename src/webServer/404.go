//Demo using the gorilla mux router. Allows routing the URL into single parameters
//e.g. /books/<bookName>/page/<pageNumber>/

package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lost, you've requested: %s\n", r.URL.Path)
}
func main() {
	//create the new router
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	//call the handleFunc on the router
	//placeholders for the dynamic content are in {}
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		//get array of the URL elements
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	//listener using the router create above
	http.ListenAndServe(":8080", r)
}
