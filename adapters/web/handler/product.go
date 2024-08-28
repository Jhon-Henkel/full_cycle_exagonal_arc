package handler

import (
	"encoding/json"
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/adapters/dto"
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, middleware *negroni.Negroni, service application.IProductService) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service))),
	).Methods("GET", "OPTIONS")
	router.Handle("/product", middleware.With(
		negroni.Wrap(createProduct(service))),
	).Methods("POST", "OPTIONS")
}

func getProduct(service application.IProductService) http.Handler {
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

func createProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDTO dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(JsonError(err.Error()))
			return
		}
		product, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(JsonError(err.Error()))
			return
		}
	})
}
