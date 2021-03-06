package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	//user:password@tcp(localhost:5555)/dbname?charset=utf8
//mydbinstance.cxcs214am8nx.us-east-1.rds.amazonaws.com
	db, err = sql.Open("mysql", "awsuser:password@tcp(mydbsinstance2.cxcs214am8nx.us-east-1.rds.amazonaws.com:3306)/test2?charset=utf8")  
//	db, err = sql.Open("mysql", "awsuser:password@tcp(mydbinstance.cxcs214am8nx.us-east-1.rds.amazonaws.com:3306)/test2?charset=utf8")                 
//	db, err = sql.Open("mysql", "awsuser:password@tcp(mydbinstance-go.cxcs214am8nx.us-east-1.rds.amazonaws.com:3306)/schema_test2?charset=utf8")                 
//	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	if err == nil { fmt.Println("no ping error, db found")}

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
