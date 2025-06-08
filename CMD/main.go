package main

import (
	"golang-app/configs"
	"golang-app/internal/authModule"
	"golang-app/internal/link"
	"golang-app/pkg/db"
	"log"
	"net/http"
)

var PORT = ":8080"

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	linkRepository := link.NewLinkRepository(db)

	authModule.AuthHandler(router, authModule.AuthDependencies{Config: conf})
	link.LinkHandler(router, link.LinkDependencies{Config: conf, LinkRepository: linkRepository})

	server := http.Server{Addr: PORT, Handler: router}

	log.Printf("Server listen and serve on port %s", PORT)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
