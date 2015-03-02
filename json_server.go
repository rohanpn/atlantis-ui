package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

type Person struct {
	Name    string   `json:name`
	Hobbies []string `json:hobbies`
}

type Application struct {

	Name         string
	Container    int8
	CPU_Shares   int8
	Memory       int32
	Dependencies []string
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "%s", p.Body)

}

func ListEnv(w http.ResponseWriter, r *http.Request) {
	var pe [7]Person
	pe[0] = Person{"john", []string{"a", "b", "c"}}
	pe[1] = Person{"john1", []string{"a1", "b1", "c1"}}
	pe[2] = Person{"john2", []string{"a2", "b2", "c2"}}
	pe[3] = Person{"john3", []string{"a2", "b2", "c2"}}
	pe[4] = Person{"john4", []string{"a2", "b2", "c2"}}
	pe[5] = Person{"john5", []string{"a2", "b2", "c2"}}
	js, err := json.Marshal(pe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func ListApp(w http.ResponseWriter, r *http.Request) {
	var App [3]Application
	App[0] = Application{"hello", 10, 2, 1024, []string{"redis", "sql", "heroku"}}
	App[1] = Application{"hello1", 10, 2, 1024, []string{"redis1", "sql1", "heroku1"}}
	App[2] = Application{"hello2", 10, 2, 1024, []string{"redis2", "sql2", "heroku2"}}
	js, err := json.Marshal(App)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
func loadPage(title string) (*Page, error) {
	filename := title // + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/view/Env", ListEnv)
	http.HandleFunc("/view/App", ListApp)

	http.HandleFunc("/view/", ViewHandler)

	fmt.Println("Listening on port 9001")
	http.ListenAndServe(":9001", nil)
}
