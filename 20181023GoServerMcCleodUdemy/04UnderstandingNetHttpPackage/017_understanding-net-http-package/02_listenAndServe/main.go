package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
	fmt.Fprintln(w, "and other lines also")
}

func main() {
	var d hotdog
	http.ListenAndServe(":9000", d)
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
