package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	templ, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	templ, err = templ.ParseFiles("two.gmao", "gopher.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	// 存在するgmaoファイルをparseした後に何が入ってるのかを見てみる
	fmt.Println("\n---------------")
	for _, t := range templ.Templates() {
		fmt.Println(*t)
	}
	fmt.Println("---------------")

	// ExecuteTemplateの第一引数はio.Writerなので、
	// os.Createの返り値を受け取ったvarhtmlをExecuteTemplateに渡せば
	// ファイルの書き込みができる。
	// os.Stdoutはただの標準出力（書き込みじゃなくてただの出力）
	varhtml, _ := os.Create("index.html")
	err = templ.ExecuteTemplate(varhtml, "gopher.gmao", nil)
	// err = templ.ExecuteTemplate(os.Stdout, "gopher.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executeで実行されるのは最初のファイルのみ
	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
