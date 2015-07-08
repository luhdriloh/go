package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.HandleFunc("/adding/", GenerateAdd)
	http.Handle("/", http.FileServer(http.Dir("pages")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	rw.Write(markdown)
}

func GenerateAdd(rw http.ResponseWriter, req *http.Request) {
	num1String := req.URL.Query().Get("number_one")
	num1, _ := strconv.ParseInt(num1String, 10, 8)

	num2String := req.URL.Query().Get("number_two")
	num2, _ := strconv.ParseInt(num2String, 10, 8)

	add := num1 + num2

	rw.Write([]byte(fmt.Sprintf("%d", add)))
}
