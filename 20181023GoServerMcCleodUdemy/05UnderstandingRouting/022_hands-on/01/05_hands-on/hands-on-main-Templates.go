package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	//"io"
)

type respo int
type respp int
type respm int

func (r respo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is alpha")
}

func (r respp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is beta, dog	")
}

func (r respm) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hi , my name is : ")
	tpl, err := template.ParseFiles("data.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "data.gohtml", "tabby")
	if err != nil {
		log.Fatalln("errors", err)
	}
}

//var tpl *template.Template

//func init() {
//	template.Must(template.ParseFiles("data.gohtml"))
//}

func main() {
	mux := http.NewServeMux()
	var a respo
	var b respp
	var me1 respm
	mux.Handle("/", a)
	mux.Handle("/dog", b)
	mux.Handle("/me", me1)
	http.ListenAndServe(":8080", mux)
}
