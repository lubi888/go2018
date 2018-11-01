package main

import (
	"log"
	"net/http"
)

func main() {
	//must have file 'index.html' or whole folder will be visible inc code
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
