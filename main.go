package main

import (
	"AdventureBook/cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
    "github.com/gorilla/mux"
)

var story cyoa.Story

func main() {
	
	r:=NewRouter()
	//port := flag.Int("port", 8000, "the localhost port")
	filename := flag.String("file", "gopher.json", "the json filefor the storybook")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	d := json.NewDecoder(f)
	if err := d.Decode(&story); err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", storyhandler)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)

	fmt.Println("starting the server at port")
	r.HandleFunc("/", storyhandler)
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}

//NewRouter .....
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}

func storyhandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if ch, ok := story[path]; ok {
		fmt.Println(ch)
		cyoa.Rendertemplate(w, "C:/Users/yashi/go/src/AdventureBook/template.html", ch)
	}
}
