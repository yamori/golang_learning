package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageDataExtra struct {
	PageTitle   string
	Todos       []Todo
	StyleString string
}

//go:embed templates/*
var templateData embed.FS

//go:embed assets/style.css
var stylCSS string

func main() {

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println(stylCSS)
	// Here I attempting to use `stylCSS` as a parameter into the template, but....
	// --> ZgotmplZ
	// go doesn't like this, identifies it as an unsafe injection
	// https://stackoverflow.com/questions/14765395/why-am-i-seeing-zgotmplz-in-my-go-html-template-output

	template, err := template.ParseFS(templateData, "templates/layout_w_assets.html")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageDataExtra{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
			StyleString: stylCSS,
		}
		template.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
