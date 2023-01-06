package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	template, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = template.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
