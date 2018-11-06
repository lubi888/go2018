package main

import (
	"fmt"
	"log"
	"net/http"
)

var c int = 0
var cString = string(c)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")

	c += 1
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "value" + string(c), //not working file downloaded  "value"+cString also not working
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	fmt.Fprintln(w, "and the values of total visits is now: ")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//var c int=0
//var cString = string(c)
//func main() {
//	http.HandleFunc("/", footsie)
//	http.HandleFunc("/bar", bar)
//	http.HandleFunc("/read", read)
//	http.Handle("/favicon.icon", http.NotFoundHandler())
//	http.ListenAndServe("8080", nil)
//}
//
//func footsie(w http.ResponseWriter, r *http.Request)  {
//	c += 1
//	http.SetCookie(w, &http.Cookie{
//		Name:  "my-cookie-visits",
//		Value: "times:" + cString,
//	})
//	fmt.Fprintln(w,"that was easy.", r.Method, "is the method")
//	fmt.Fprintln(w,"has been visited: ", c ,"times")
//}
//
//func bar(w http.ResponseWriter, r *http.Request){
//	fmt.Println("that was also bar ez")
//}
//
//func read(w http.ResponseWriter, r *http.Request) {
//	//if (http.Cookie{"my-cookie-visits"})
//	c1, err:= r.Cookie("my-cookie-visits")
//	if err != nil {
//		log.Fatalln(err)
//		log.Println(err)
//	}
//	fmt.Fprintln(w, "what we read:", c1.Name, c1.Value)
//}
