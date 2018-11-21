package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting server now in main. open web browser  1")
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
	//fmt.Println("server started and fed in main  3")   //never reached
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("server started in foo   2")
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
