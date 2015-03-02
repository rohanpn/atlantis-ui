package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<div>%s</div>", p.Body)

	fmt.Println("env list is : ", listEnvironment())
}

func loadPage(title string) (*Page, error) {
	filename := title // + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func listEnvironment() []string {
	envList := []string{"env1", "env2", "env3"}
	return envList

}

func main() {
	http.HandleFunc("/view/", viewHandler)

	fmt.Println("listening on 9001")
	http.ListenAndServe(":9001", nil)
}
