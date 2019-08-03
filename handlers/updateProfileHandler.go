package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/teohrt/cruddyAPI/entity"
	"github.com/teohrt/cruddyAPI/service"
)

// TODO
func UpdateProfileHandler(svc service.Service) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		logger := zerolog.Ctx(r.Context())
		decoder := json.NewDecoder(r.Body)

		profile := new(entity.Profile)
		if err := decoder.Decode(profile); err != nil {
			logger.Debug().Err(err).Msg("Bad req body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "TODO")
	}
}
