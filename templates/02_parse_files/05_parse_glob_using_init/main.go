package main

import (
	"log"
	"os"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// Executeで実行されるのは最初に読み込んだファイルのみ
	err := templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "wgopher.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
