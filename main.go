package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Connect4 struct {
	Title string
	Col   []int
	Row   []int
}

var mainPage = Connect4{
	Title: "Connect4",
	Col:   []int{0, 1, 2, 3, 4, 5, 6},
	Row:   []int{5, 4, 3, 2, 1, 0},
}

func main() {
	// mux := http.NewServeMux()
	http.HandleFunc("/", handleRoot)
	http.Handle("/template_files/css/", http.StripPrefix("/template_files/css/", http.FileServer(http.Dir("./template_files/css/"))))
	http.Handle("/template_files/js/", http.StripPrefix("/template_files/js/", http.FileServer(http.Dir("./template_files/js/"))))

	fmt.Println("server listening to port :8888")
	var err = http.ListenAndServe(":8888", nil)

	fmt.Println("ended with: ", err)
}

func handleRoot(
	w http.ResponseWriter,
	r *http.Request,
) {
	// fmt.Fprintf(w, "Hello World")
	t, _ := template.ParseFiles("template_files/gohtml/index.go.html")
	var err = t.Execute(w, mainPage)
	fmt.Println(err)
}
