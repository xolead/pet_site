package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	port = "3000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/info", infoPage)
	mux.HandleFunc("/create", createPage)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	log.Println("Запуск сервера")

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
