package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//split 1 txt into 3
	t:= `toddy`
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		`
	u := `</body>
		</html>`
	w.Write([]byte(s))
	w.Write([]byte(t))
	w.Write([]byte(u))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
