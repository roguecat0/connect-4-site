package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
)

type Connect4 struct {
	Title string
	Col   []int
	Row   []int
}

type Response struct {
	CpuMove       int `json:"cpu_move"`
	GameStatePre  int `json:"game_state_pre"`
	GameStatePost int `json:"game_state_post"`
}
type Request struct {
}

var mainPage = Connect4{
	Title: "Connect4",
	Col:   []int{0, 1, 2, 3, 4, 5, 6},
	Row:   []int{5, 4, 3, 2, 1, 0},
}

func main() {
	// mux := http.NewServeMux()
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/connect4", handleRoot)
	http.HandleFunc("GET /move/{moves}", handleGet)
	http.Handle("/template_files/css/", http.StripPrefix("/template_files/css/", http.FileServer(http.Dir("./template_files/css/"))))
	http.Handle("/template_files/js/", http.StripPrefix("/template_files/js/", http.FileServer(http.Dir("./template_files/js/"))))

	fmt.Println("server listening to port :8082")
	var err = http.ListenAndServe("0.0.0.0:8082", nil)

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

func handleGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	var moves = r.PathValue("moves")
	fmt.Println(moves)
	var res = Response{
		CpuMove:       rand.IntN(5),
		GameStatePre:  2,
		GameStatePost: 2,
	}
	json.NewEncoder(w).Encode(res)
}
