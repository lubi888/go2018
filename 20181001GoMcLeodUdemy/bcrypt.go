package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	//use go get "golang.org/x/crypto/bcrypt" to import
)

func main() {
	s := `password123`
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("all smoothe")
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password1234`
	loginPword12 := `password123`

	//func CompareHashAndPassword(hashedPassword, password []byte) error
	//cant redeclare 'err' so just use = instead of assignment :=
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		//return  //this ends prog?
	}
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword12))
	if err != nil {
		fmt.Println("failed again. no login")
	} else {
		fmt.Println("You're logged in, pwd hashs matched")
	}
}
