package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var suitU = [4]string{"\u2660", "\u2665", "\u2663", "\u2666"}
var suitUW = [4]string{"&#9824;", "&#9829;", "&#9827;", "&#9830;"}
var deckU = [3]string{"\U0001f0a1", "\U0001F0D1", "\U0001f0b8"}
var number = [13]string{"ace", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ", " 10 ", "jack", "queen", "king"}
var tpl *template.Template
var deckR = make(map[string]string)

type cardface struct {
	suit   string
	number string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fmt.Println("setting server now")
	http.HandleFunc("/", deck)
	http.HandleFunc("/num", num)
	http.HandleFunc("/card", card)
	http.HandleFunc("/deck", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", suitU)
}

func num(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", number)
}

func card(w http.ResponseWriter, r *http.Request) {
	//not working
	var cf cardface
	cf = cardface{"ace", "diamonds"}
	tpl.ExecuteTemplate(w, "index.gohtml", cf)
}

//func deck, see file deck-unicode.go

//func deck(w http.ResponseWriter, r *http.Request) {
//	//not working
//	//tpl.ExecuteTemplate(w, "index.gohtml", deckR["ace"])
//	tpl.ExecuteTemplate(w, "index.gohtml", a)
//}
