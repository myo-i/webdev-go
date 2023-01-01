package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	// templates配下のファイル全てを同時に読み込むことも可能
	templ, err := template.ParseGlob("templates/*")
	// templ, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Executeで実行されるのは最初のファイルのみ
	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "gopher.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// templates/*.gohtmlで.gohtmlのみを読み込んだ場合、下記で呼び出せなくなる
	err = templ.ExecuteTemplate(os.Stdout, "hello.html", nil)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
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
