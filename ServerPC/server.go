// hello60
package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>你好!</h1>")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:5000", h)
}