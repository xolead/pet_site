package main

import (
	//"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/registration", regPage)
	mux.HandleFunc("/create", createPage)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	log.Println("Запуск сервера")

	err := http.ListenAndServe(":80", mux)
	log.Fatal(err)
}
