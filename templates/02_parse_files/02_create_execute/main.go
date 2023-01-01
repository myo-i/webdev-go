package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templ, err := template.ParseFiles("templ.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	err = templ.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
