package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/")))

	err := http.ListenAndServe(":3141", nil)
	if err != nil {
		panic(err)
	}
}
