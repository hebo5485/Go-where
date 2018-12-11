package main

import (
	. "fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Println(r.Form)
	Println("path", r.URL.Path)
	Println("scheme", r.URL.Scheme)
	Println(r.Form["url_long"])
	for k, v := range r.Form {
		Println("KEY:", k)
		Println("val:", strings.Join(v, ""))
	}
	Fprintf(w, "Hey buddy")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
