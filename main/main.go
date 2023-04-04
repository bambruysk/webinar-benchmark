package main

import (
	"html/template"
	"log"
	"net/http"
	"net/http/pprof"
)

var tmpl = template.Must(template.ParseFiles("template.html"))

type DataHTML struct {
	Title   string
	Header  string
	Content string
}

func handleHTML(w http.ResponseWriter, r *http.Request) {
	data := &DataHTML{
		Title:   "My Page",
		Header:  "Welcome to my page!",
		Content: "This is the content of my page.",
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handleHTML)
	http.HandleFunc("/json", handleJSON)

	http.Handle("cpu", pprof.Handler("sevice"))

	log.Fatal(http.ListenAndServe("127.0.0.1:9080", nil))
}
