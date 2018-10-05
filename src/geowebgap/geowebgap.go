package geowebgap

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Accept Address</title>
      </head>
      <body>
        <h1>Accept Address</h1>
        <p>Please enter your address:</p>
        <form action="" method="post" accept-charset="utf-8">
	  <input type="text" name="str" value="Type address..." id="str">
	  <input type="submit" value=".. and see the image!">
        </form>
      </body>
    </html>
`
