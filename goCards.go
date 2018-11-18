package main
//import "unicode"
import "fmt"
import "math/rand"
import "strings"
//var suit [4]string{hearts, diamonds, clubs, spades}
//suit = append(suit, "hearts", "diamonds", "trumps", "spades")
//{hearts, diamonds, clubs, spades}
//var number [13]string {"1 ace", "2 two", "3 three", "4 four", "5 five", "6 six", "7 seven", "8 eight", "9 nine", "10 ten", "11 Jack", "12 Queen", "13 King"}
func main() {
	suit := [4]string{"hearts", "diamonds", "clubs", "spades"}
	suitU := [4]string{"\u2660", "\u2665", "\u2663", "\u2666"}
	deckU := [3]string{"\U0001f0a1", "\u1F0D1", "\U0001f0b8"}   //1F0A0—1F0FF
	deckR := make(map[string]string)
	//deckRnum := [52]int{}
	number := [13]string{"ace", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ", " 8 ", " 9 ", " 10 ", "jack", "queen", "king"}

	fmt.Println(suit[1])
	fmt.Println(number[1])
	fmt.Println(number[12])
	fmt.Print("the", number[11], "of", suit[2])
	fmt.Println("--------------")
		for i, s := range suitU {
			for p, t := range number {
			deckR[s] += t
				fmt.Println(i+1, "\tdecrR increased", p+1)
			}
		}
	fmt.Println(deckU[0])
	fmt.Println(deckU[2])

	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
	fmt.Println(deckR)
	fmt.Println(deckR["queen" ])
	}
