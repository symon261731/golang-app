package req

import (
	"golang-app/pkg/res"
	"log"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, request *http.Request) (*T, error) {

	body, decodeError := Decode[T](request.Body)
	if decodeError != nil {
		log.Printf("Error decoding json: %v", decodeError)
		res.Json(w, decodeError.Error(), 400)
		return nil, decodeError
	}

	valError := IsValid(body)
	if valError != nil {
		res.Json(w, valError.Error(), 400)
		return nil, valError
	}

	return &body, nil
}
