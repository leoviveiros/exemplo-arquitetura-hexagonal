package handler

import (
	"encoding/json"
	"go-hexagonal/adapters/dto"
	"go-hexagonal/application"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewProductHandler(router *mux.Router, middleware *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/product", middleware.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/product/{id}/enable", middleware.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")

	router.Handle("/product/{id}/disable", middleware.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
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

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDTO dto.Product

		err := json.NewDecoder(r.Body).Decode(&productDTO)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDTO.Name, productDTO.Price)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Enable(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Disable(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}