package main

import (
	"../../webtool"
	"./app/articles"
	"fmt"
	"html"
	"net/http"
	"os"
)

func main() {

	basePath := "/articles/"
	router := util.NewHttpRouter().SetBasePath(basePath)
	router.SetRoute("", articles.TopPage)
	router.SetRoute("article/<id>", articles.ArticlePage)

	// for default error
	router.SetError("default", func(message string, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Error, %q\n", html.EscapeString(message))
	})
	// for 404 error
	router.SetError("404", func(message string, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "404 Error, %q\n", html.EscapeString(message))
	})
	// for other error
	router.SetError("other", func(message string, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "other Error, %q\n", html.EscapeString(message))
	})

	http.Handle(basePath, router)
	fmt.Fprint(os.Stdout, "http://localhost:8000/articles/\n")
	fmt.Fprint(os.Stdout, "http://localhost:8000/articles/article/1\n")
	err := http.ListenAndServe("localhost:8000", nil)
	// err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
