package main

import "fmt"

//var suit [4]string{hearts, diamonds, clubs, spades}

//suit = append(suit, "hearts", "diamonds", "trumps", "spades")

//{hearts, diamonds, clubs, spades}
//var number [13]string {"1 ace", "2 two", "3 three", "4 four", "5 five", "6 six", "7 seven", "8 eight", "9 nine", "10 ten", "11 Jack", "12 Queen", "13 King"}

func main() {
	suit := [4]string{"hearts", "diamonds", "clubs", "spades"}

	number := [13]string{"1 ace", "2 two", "3 three", "4 four", "5 five", "6 six", "7 seven", "8 eight", "9 nine", "10 ten", "11 Jack", "12 Queen", "13 King"}

	fmt.Println(suit[1])
	fmt.Println(number[1])
	fmt.Println(number[12])
	fmt.Print("the", number[11], "of", suit[2])
	fmt.Println("--------------")
	for i, s := range suit {
		fmt.Println(i, "and the suit range is", s)
		for t, r := range number {
			fmt.Println(t+1, "while the card is", r, "of", s)
		}
	}
}
