package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("ui/html/info.html")
	if err != nil {
		log.Println((err.Error()))
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Serve Error", 500)
	}
}

func createPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Get NO"))
		return
	}
	w.Write([]byte("LOX"))
}
