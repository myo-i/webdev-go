package main

import (
	"log"
	"os"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("templ.gohtml"))
}

func main() {
	err := templ.ExecuteTemplate(os.Stdout, "templ.gohtml", `Ramen, Black, and 2`)
	if err != nil {
		log.Fatalln(err)
	}
}
