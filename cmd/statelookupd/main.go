package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gitlab.com/LanceH/state-lookup/geom"
)

var states = make(map[string]geom.State)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", search)
	http.ListenAndServe(":8080", nil)
}

func init() {
	b, err := ioutil.ReadFile("../../data/states.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(b, &states)
	if err != nil {
		log.Panic(err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lat, _ := r.Form["lat"]
	lng, _ := r.Form["lng"]
	if len(lat) < 1 || len(lng) < 1 {
		return
	}
	x, _ := strconv.ParseFloat(lat[0], 64)
	y, _ := strconv.ParseFloat(lng[0], 64)
	w.Write([]byte(fmt.Sprint(x, y)))
}
