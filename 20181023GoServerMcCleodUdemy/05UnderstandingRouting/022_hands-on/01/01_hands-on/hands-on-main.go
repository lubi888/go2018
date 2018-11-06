package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is alpha")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is beta, dog	")
}

func me1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hi , my name is : ")
}
func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", me1)
	http.ListenAndServe(":8080", nil)
}
