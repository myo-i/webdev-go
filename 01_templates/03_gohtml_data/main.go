package main

import (
	"log"
	"os"
	"strconv"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("templ.gohtml"))
}

func main() {
	// ExecuteTemplateの第三引数にデータを与えれば、templateにデータを渡せる。
	// sliceなどを用いて複数渡す場合はReadmeに記載
	err := templ.ExecuteTemplate(os.Stdout, "templ.gohtml", "sushi"+strconv.Itoa(33))
	if err != nil {
		log.Fatalln(err)
	}
}
