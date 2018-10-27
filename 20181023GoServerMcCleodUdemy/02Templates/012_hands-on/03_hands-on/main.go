package main

import (
	"text/template"
	"os"
	"log"
)

type hotel struct {
	Name, Address, City, Region string
	Zip int
}

type region struct {
	hotels []hotel
	Northern, Central, Southern string
}

type caliornia struct {
	regions []region
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	///californias := []caliornia{

		err := tpl.Execute(os.Stdout, californias)
	if err != nil {
		log.Fataln(err)
	}

	}


