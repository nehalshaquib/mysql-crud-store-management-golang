package server

import (
	"encoding/json"
	"errors"
	"golang-store-management/pkg/db"
	"golang-store-management/pkg/model"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func (S *Store) GetItems(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	items, err := db.AllItems(S.DB)
	if err != nil {
		logrus.Errorln(err)
	}

	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		logrus.Errorln(err)
	}
}

func (S *Store) CreateNewItem(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	newItem := model.Item{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newItem)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateCreateItem(newItem)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateItem(S.DB, newItem)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			http.Error(w, "Item ID already exists, please enter a new unique id", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.SuccessResponse{Detail: "Item created successfully"})
}

func validateCreateItem(item model.Item) error {
	if item.ID == 0 {
		return errors.New("id field empty or missing")
	}
	if item.Name == "" {
		return errors.New("name field empty or missing")
	}
	if item.Price == 0 {
		return errors.New("price field empty or missing")
	}
	if item.Type == "" {
		return errors.New("type field empty or missing")
	}
	if item.Quantity == 0 {
		return errors.New("quantity field empty or missing")
	}

	return nil
}
