package router

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"github.com/gorilla/mux"
)

const APIVersion = 1

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	routerVersionGroup := router.PathPrefix(fmt.Sprintf("/v%v/", APIVersion)).Subrouter()

	routerVersionGroup.
		Methods("GET").
		Path("/").
		Name("Index").
		Handler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
		
				json.NewEncoder(w).Encode(map[string]string{
					"docker_version": os.Getenv("VERSION"),
					"environment": os.Getenv("APP_ENV"),
				})
			}))

	return router
}
