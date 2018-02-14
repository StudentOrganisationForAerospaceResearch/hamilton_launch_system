package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func absPath(relativePath string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err, true)
	result := path.Join(dir, relativePath)
	return result
}

func check(err error, exit bool) {
	if exit {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "%v", r)
				os.Exit(1)
			}
		}()
		if err != nil {
			panic(err)
		}
	} else {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	}
}

func genericPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL was ", r.URL)
	var t *template.Template
	if r.URL.String() == "/" {
		t = template.Must(template.ParseFiles(
			absPath("templates/base.html"),
			absPath("templates/pages/home.html")))
	} else {
		page := absPath("templates/pages/" + r.URL.String() + ".html")
		if _, err := os.Stat(page); os.IsNotExist(err) {
			// page does not exist
			page = "templates/pages/404.html"
		}
		t = template.Must(template.ParseFiles(
			absPath("templates/base.html"), page))
	}
	err := t.ExecuteTemplate(w, "base", nil)
	check(err, true)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api" {
		http.NotFound(w, r)
		return
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, absPath("resources/images/favicon.ico"))
}

func cache(h http.Handler) http.Handler {
	var cacheHeaders = map[string]string{
	// "Cache-Control": "public, max-age=2592000",
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range cacheHeaders {
			w.Header().Set(k, v)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {
	http.HandleFunc("/", genericPageHandler)
	http.HandleFunc("/api", apiHandler)
	http.Handle("/resources/", cache(http.StripPrefix("/resources/", http.FileServer(http.Dir(absPath("resources"))))))
	http.HandleFunc("/favicon.ico", faviconHandler)

	port := "8000"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	err := http.ListenAndServe(":"+port, nil)
	check(err, true)
}
