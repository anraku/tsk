package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:password@tcp(0.0.0.0:3306)/tsk?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return db, nil
}
