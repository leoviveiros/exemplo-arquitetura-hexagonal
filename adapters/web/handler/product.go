package handler

import (
	"encoding/json"
	"go-hexagonal/application"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewProductHandler(router *mux.Router, middleware *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}