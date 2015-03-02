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

type Environment struct {
	Name         string
	Container    int32
	CPU_Shares   int32
	Memory       int32
	Dependencies []string
	Sha          []int16
}

type Application struct {
	Name string
	Env  []Environment
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
	//var App []Application = {{"app1",[]Environment{{"env1",10,2,1024,[]string{"a","b","c"},[]int{1,2,3}}}}}
	//App[1]={"app1",[]Environment{{"env1",10,2,1024,[]string{"a","b","c"},[]int{1,2,3}}}}
	App[0] = Application{"app1", []Environment{{"env1", 1, 2, 3, []string{"redis", "sql", "heroku"}, []int16{10, 11, 12}},
		{"env12", 4, 5, 6, []string{"redis1", "sql1", "heroku1"}, []int16{101, 102, 103}}}}

	App[1] = Application{"app2", []Environment{{"env2", 1, 2, 3, []string{"redis2", "sql2", "heroku2"}, []int16{20, 21, 22}},
		{"env21", 4, 5, 6, []string{"redis21", "sql22", "heroku23"}, []int16{201, 202, 203}}}}

	App[2] = Application{"app3", []Environment{{"env3", 1, 2, 3, []string{"redis3", "sql3", "heroku3"}, []int16{30, 31, 32}},
		{"env31", 4, 5, 6, []string{"redis31", "sql32", "heroku33"}, []int16{301, 302, 303}}}}

	js, err := json.Marshal(App)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func loadPage(title string) (*Page, error) {
	filename := title
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

	fmt.Println("Listening on port 9002")
	http.ListenAndServe(":9002", nil)
}
