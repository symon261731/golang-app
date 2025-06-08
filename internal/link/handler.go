package link

import (
	"golang-app/configs"
	"golang-app/pkg/req"
	"golang-app/pkg/res"
	"log"
	"net/http"
)

type LinkDependencies struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}

type LinkModule struct {
	LinkDependencies
}

func LinkHandler(router *http.ServeMux, dependecies LinkDependencies) {
	linkInstance := &LinkModule{
		LinkDependencies: LinkDependencies{
			LinkRepository: dependecies.LinkRepository,
			Config:         dependecies.Config,
		},
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

func (handler *LinkModule) CreateLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	body, err := req.HandleBody[RequestLinkData](w, r)
	if err != nil {
		return
	}

	generatedLink := NewLink(body.Url)

	dbErr := handler.LinkRepository.Create(generatedLink)

	if dbErr != nil {
		log.Println(dbErr)
		w.WriteHeader(500)
		return
	}

	res.Json(w, generatedLink, 200)
}

func (handler *LinkModule) UpdateLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.WriteHeader(405)
		return
	}

	id := r.PathValue("id")
	log.Println(id)

}

func (handler *LinkModule) DeleteLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		return
	}

	id := r.PathValue("id")

	dbActionErr := handler.LinkRepository.Delete(id)

	if dbActionErr != nil {
		log.Println(dbActionErr)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}
