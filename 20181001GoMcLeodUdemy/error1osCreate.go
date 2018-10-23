package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("namesBond7.txt") //creates file in directory above this one
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)
}

//package main
//
//import (
//"fmt"
//"io/ioutil"
//"os"
//)
//
//func main() {
//f, err := os.Open("namesBond7.txt")
//if err != nil {
//fmt.Println(err)
//return
//}
//defer f.Close()
//
//bs, err := ioutil.ReadAll(f)
//if err != nil {
//fmt.Println(err)
//return
//}
//
//fmt.Println(string(bs))
//}
