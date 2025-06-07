package authModule

import (
	"fmt"
	"golang-app/configs"
	"golang-app/pkg/req"
	"golang-app/pkg/res"
	"log"
	"net/http"
)

type AuthDependencies struct {
	*configs.Config
}

type AuthModule struct {
	AuthDependencies
}

func AuthHandler(router *http.ServeMux, dependencies AuthDependencies) *AuthModule {
	handler := &AuthModule{
		dependencies,
	}
	router.HandleFunc("/auth/login", handler.Authentication)
	router.HandleFunc("/auth/register", handler.Registration)

	return handler
}

func (authModule *AuthModule) Authentication(w http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	body, err := req.HandleBody[LoginRequest](w, request)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(body)

	resData := LoginResponse{
		Token: "123",
	}

	res.Json(w, resData, 201)
	log.Print("Authentication")
}

func (authModule *AuthModule) Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	log.Print("Registration")
}
