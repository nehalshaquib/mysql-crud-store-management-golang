package main

import (
	"fmt"
	"golang-store-management/pkg/db"
	"golang-store-management/pkg/server"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	store := server.Store{}
	db, err := db.CreateDB()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	store.DB = db

	r := store.CreateRoutes()
	store.Server = &http.Server{Addr: ":" + "8090", Handler: r}
	fmt.Println("Starting Server...")
	// go func() {
	err = store.Server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	// }()
	fmt.Println("Server Started")
}
