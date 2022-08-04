package server

import (
	"fmt"
	"golang-store-management/pkg/model"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Store struct {
	model.Item
	DB     *gorm.DB
	Server *http.Server
}

func (S *Store) CreateRoutes() *mux.Router {
	fmt.Println("in CreateRoutes")
	router := mux.NewRouter()
	router.Use(auth)
	router = router.PathPrefix("/store").Subrouter()
	router.HandleFunc("/getItems", S.GetItems).Methods(http.MethodGet)
	router.HandleFunc("/createItem", S.CreateItem).Methods(http.MethodPost)

	return router
}

func auth(next http.Handler) http.Handler {
	fmt.Println("in auth")
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
		}
		if token != "1234abcd" {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
	return http.HandlerFunc(fn)
}
