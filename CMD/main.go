package main

import (
	"golang-app/configs"
	"golang-app/pkg/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		return
	})
	server := http.Server{Addr: ":8080", Handler: router}

	log.Println("Server listen and serve")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
