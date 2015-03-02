package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//	"os"
	"log"
)

type Uptime struct {
	IP        string
	TimeStamp string
	LoadAvg   float32
	UpSince   int64
}

func upTime(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("data is : ", w)

	err := ioutil.WriteFile("timelog.txt", body, 0660)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	http.HandleFunc("/uptime", upTime)

	fmt.Println("Listening on port 9000")
	http.ListenAndServe(":9000", nil)

}
