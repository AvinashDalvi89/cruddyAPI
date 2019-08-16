package app

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"

	"github.com/teohrt/cruddyAPI/dbclient"
	"github.com/teohrt/cruddyAPI/handlers"
	"github.com/teohrt/cruddyAPI/service"
)

func Start() {
	// PORT := os.Getenv("SERVER_PORT")
	config := dbclient.Config{}
	env.Parse(&config)
	svc := service.New(&config)
	v := validator.New()

	router := mux.NewRouter()

	subRouter := router.PathPrefix("/health").Subrouter()
	subRouter.HandleFunc("/ping", handlers.Health()).Methods(http.MethodGet)

	subRouter = router.PathPrefix("/cruddyAPI").Subrouter()
	subRouter.HandleFunc("/profiles", handlers.CreateProfile(svc, v)).Methods(http.MethodPost)
	subRouter.HandleFunc("/profiles/{id}", handlers.GetProfile(svc)).Methods(http.MethodGet)
	subRouter.HandleFunc("/profiles/{id}", handlers.UpdateProfile(svc, v)).Methods(http.MethodPut)
	subRouter.HandleFunc("/profiles/{id}", handlers.DeleteProfile(svc)).Methods(http.MethodDelete)

	// fmt.Println("Server running locally and listening on port :" + PORT)
	// log.Fatal(http.ListenAndServe(":"+PORT, router))

	log.Fatal(gateway.ListenAndServe("", router))
}
