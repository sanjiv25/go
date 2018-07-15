package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {

		rw.Write([]byte("Hello Go"))

	})

	http.ListenAndServe(":8000", nil)
}
