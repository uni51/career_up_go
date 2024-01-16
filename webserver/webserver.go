package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "レッツゴー\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
