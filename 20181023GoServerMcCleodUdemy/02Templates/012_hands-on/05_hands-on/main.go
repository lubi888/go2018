package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast, Lunch, Dinner string
}
type menus []menu

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("menu.gohtml"))
}
func main() {
	m1 := menus{
		{
			Breakfast: "cereal bacon",
			Lunch:     "sammy",
			Dinner:    "seafood",
		},
		{
			Breakfast: "sausages",
			Lunch:     "soup",
			Dinner:    "Pie",
		},
	}

	err := tpl.Execute(os.Stdout, m1)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}
