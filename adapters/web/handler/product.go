package handler

import (
	"encoding/json"
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, middleware *negroni.Negroni, service application.IProductService) {
	router.Handle("/product/{id}", middleware.With(
		negroni.Wrap(getProduct(service))),
	).Methods("GET", "OPTIONS")
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
