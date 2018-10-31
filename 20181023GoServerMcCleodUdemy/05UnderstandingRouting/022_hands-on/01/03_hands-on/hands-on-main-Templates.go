package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	//"io"
)

func a(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "this is alpha")
}

func d(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "This is beta, dog	")
}

func me1(w http.ResponseWriter, req *http.Request)  {
	//io.WriteString(w, "hi , my name is : ")
	tpl, err:= template.ParseFiles("data.gohtml")
	if err != nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w,"data.gohtml", "tabby")
	if err != nil {
		log.Fatalln("errors", err)
	}
}

//var tpl *template.Template


//func init() {
//	template.Must(template.ParseFiles("data.gohtml"))
//}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", me1)
	http.ListenAndServe(":8080", nil)
}
