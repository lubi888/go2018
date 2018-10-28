package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast, Lunch, Dinner string
}
type restaurant struct {
	ResName string
	Menus   []menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("menu.gohtml"))
}

func main() {
	rs1 := restaurants{
		restaurant{
			ResName: "snor",
			Menus: []menu{
				menu{
					Breakfast: "tuna",
					Lunch:     "sandwiches",
					Dinner:    "burgers",
				}}},
		restaurant{
			ResName: "totty",
			Menus: []menu{
				menu{Breakfast: "bacon",
					Lunch:  "soup",
					Dinner: "sushi",
				}},
		},
	}
	err := tpl.Execute(os.Stdout, rs1)
	//err := tpl.Execute(os.Stdout, m1)

	if err != nil {
		log.Fatalln(err)
	}
}
