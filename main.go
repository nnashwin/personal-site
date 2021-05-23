package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const staticDirDefault = "./static/"
const staticFlagUsage = "path to static directory"
const templateDirDefault = "templates"
const templateDirUsage = "path to template directory"

// use the closure to store the templateDir string
// return a function that maps the writer interface in order to be applicable
// to serve in HandleFunc
func serveTemplate(templateDir string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cleanPath := filepath.Clean(r.URL.Path)

		// handles sitemap.xml and robots.txt in the root directory
		if strings.Contains(cleanPath, ".txt") || strings.Contains(cleanPath, ".xml") {
			http.ServeFile(w, r, "static/"+cleanPath)
			return
		}

		fPath := filepath.Join(templateDir, cleanPath+".html")
		if cleanPath == "/" {
			fPath = filepath.Join(templateDir, "home.html")
		}

		if strings.Contains(fPath, "draft") {
			http.Redirect(w, r, r.URL.Host, http.StatusMovedPermanently)
			return
		}

		lPath := filepath.Join(templateDir, "layout.html")

		info, err := os.Stat(fPath)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
		}

		if info.IsDir() {
			http.NotFound(w, r)
			return
		}

		t, err := template.ParseFiles(lPath, fPath)
		if err != nil {
			fmt.Println("template parsing err: ", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = t.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			fmt.Println("template execution err: ", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Printf("serving %s\n", cleanPath)
	}
}

func main() {
	var staticDir string
	var templateDir string

	flag.StringVar(&staticDir, "static", staticDirDefault, staticFlagUsage)
	flag.StringVar(&staticDir, "sd", staticDirDefault, staticFlagUsage)

	flag.StringVar(&templateDir, "templates", templateDirDefault, templateDirUsage)
	flag.StringVar(&templateDir, "td", templateDirDefault, templateDirUsage)

	flag.Parse()

	fs := http.FileServer(http.Dir(staticDir))
	// strip the static prefix from the api in order to make more memorable paths
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// add the templateDir string
	http.HandleFunc("/", serveTemplate(templateDir))
	fmt.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
