package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//switch req.a {
	fmt.Println(req.Header, ":= this req.Header\n")
	fmt.Println(req.Body, ":= try the req.Body\n")
	fmt.Println(req.URL, ":= is the req.url\n")
	fmt.Println(req.URL.Host,":= is the url.host\n")

	io.WriteString(w, "?:= req.Header\n" +
		"map[Cache-Control:[max-age=0] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36] Dnt:[1] Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8] Accept-Encoding:[gzip, deflate, br] Accept-Language:[en-GB,en;q=0.9,en-US;q=0.8,de;q=0.7] Connection:[keep-alive]]")

	io.WriteString(w, "\n/favicon.ico := is the req.url")

	switch req.URL.Path {

	case "/dog":
		io.WriteString(w, "\ndoggy doggy doggy" +
			"\nis in the req.URL.Path /doggo")

	case "/cat":
		io.WriteString(w, "\nkitty kitty kitty" +
			"\nis in the req.URL.Path /cat")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
