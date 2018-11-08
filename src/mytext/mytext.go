package mytext

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc( w http.ResponseWriter, r *http.Request) {
		fmt.print(w, "hello. this is my first go web app for google app engine")
	}
