package main

import (
	"fmt"
	"time"

	"github.com/delp/hungy/internal/hungy"

	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

func open(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("additem.html")
	t.Execute(w, nil)
}

func additem(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("deck.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	item := new(hungy.Item)
	item.Description = r.FormValue("description")
	//date := r.FormValue("expires")
	item.Acquired = time.Now()
	//fmt.Println(item.Acquired)

	//layout := "2006-01-02T15:04:05.000Z"
	//str := fmt.Sprintf("%sT%sZ", r.FormValue("date"), r.FormValue("time"))
	//t, err := time.Parse(layout, str)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(t)

	b, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f.Write(b)
	f.Close()
	fmt.Println("item added")
}

func main() {
	http.HandleFunc("/additem", additem)
	http.HandleFunc("/", open)
	http.ListenAndServe(":8080", nil)

}
