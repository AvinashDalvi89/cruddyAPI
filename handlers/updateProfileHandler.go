package handlers

import (
	"fmt"
	"net/http"

	"github.com/teohrt/cruddyAPI/service"
)

// TODO
func UpdateProfileHandler(svc service.Service) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TODO")
	}
}
