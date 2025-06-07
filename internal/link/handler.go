package link

import (
	"golang-app/configs"
	"golang-app/pkg/req"
	"log"
	"net/http"
)

type LinkDependencies struct {
	*configs.Config
}

type LinkModule struct {
	LinkDependencies
}

func LinkHandler(router *http.ServeMux, dependecies LinkDependencies) {
	linkInstance := &LinkModule{
		dependecies,
	}

	router.HandleFunc("/link/create", func(writer http.ResponseWriter, request *http.Request) {
		linkInstance.CreateLink(writer, request)
	})
	router.HandleFunc("/link/update/{id}", func(writer http.ResponseWriter, request *http.Request) {
		linkInstance.UpdateLink(writer, request)
	})
	router.HandleFunc("/link/remove/{id}", func(writer http.ResponseWriter, request *http.Request) {
		linkInstance.DeleteLink(writer, request)
	})
}

func (*LinkModule) CreateLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	body, err := req.HandleBody[struct {
		Url string `json:"url" validate:"required"`
	}](w, r)

	if err != nil {

		return
	}
	generatedLink := NewLink(body.Url)

	log.Print(generatedLink)
	w.WriteHeader(200)

}

func (*LinkModule) UpdateLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.WriteHeader(405)
		return
	}

	id := r.PathValue("id")

	log.Println(id)

}

func (*LinkModule) DeleteLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		return
	}
}
