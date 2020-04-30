package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const staticFilePath = "./static/"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lPath := filepath.Join("templates", "layout.html")
	fPath := filepath.Join("templates", "home.html")

	t, err := template.ParseFiles(lPath, fPath)
	if err != nil {
		fmt.Println("template parsing err: ", err)
	}

	err = t.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		fmt.Println("template execution err: ", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homeHandler)
	fmt.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
