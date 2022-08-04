package db

import (
	"fmt"
	"golang-store-management/pkg/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, error) {
	fmt.Println("in CreateDB DB")
	db, err := gorm.Open(sqlite.Open("store.db"))
	if err != nil {
		fmt.Println(err)
		logrus.Errorln(err)
		return db, err
	}
	db.AutoMigrate(&model.Item{})

	return db, nil
}

func AllItems(db *gorm.DB) ([]model.Item, error) {
	fmt.Println("in FetchItems DB")
	items := []model.Item{}
	rows, err := db.Raw(`select * from items;`).Rows()
	if err != nil {
		logrus.Errorln(err)
		return items, err
	}
	for rows.Next() {
		db.ScanRows(rows, &items)
	}
	fmt.Println("Items:", items)
	return items, nil
}

func CreateItem(db *gorm.DB, item model.Item) {
	db.Create(&item)
}
