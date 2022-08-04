package server

import (
	"encoding/json"
	"fmt"
	"golang-store-management/pkg/db"
	"golang-store-management/pkg/model"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (S *Store) GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetItems")
	items, err := db.AllItems(S.DB)
	if err != nil {
		logrus.Errorln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		logrus.Errorln(err)
	}
}

func (S *Store) CreateItem(w http.ResponseWriter, r *http.Request) {
	newItem := model.Item{}
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		logrus.Errorln(err)
	}
	db.CreateItem(S.DB, newItem)
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Item Created")
}
