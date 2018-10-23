package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("namesBond7.txt")
	if err != nil {
		fmt.Println(err, "is err 1")
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err, "is err 2")
		return
	}

	fmt.Println(string(bs))
}
